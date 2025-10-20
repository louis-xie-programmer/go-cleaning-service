package service

import (
    "bytes"
    "encoding/json"
    "fmt"
    "os"
    "regexp"

    "golang.org/x/net/html"
    "github.com/antchfx/htmlquery"
    "go-cleaning-service/model"
)

// 定义规则结构
type FieldRule struct {
    XPath string `json:"xpath"`
    Regex string `json:"regex,omitempty"`
}

type Rule struct {
    RootXPath  string               `json:"root_xpath"`
    GroupXPath string               `json:"group_xpath,omitempty"`
    Fields     map[string]FieldRule `json:"fields"`
}

type RulesConfig struct {
    Rules map[string]Rule `json:"rules"`
}

var ruleConfig RulesConfig

// LoadRules 从文件加载规则配置
func LoadRules(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, &ruleConfig)
}

// CleanHTMLData 对 HTML 原文进行清洗
func CleanHTMLData(source string, htmlContent []byte) ([]model.CleanData, error) {
    rule, ok := ruleConfig.Rules[source]
    if !ok {
        return nil, fmt.Errorf("no rule found for source: %s", source)
    }

    doc, err := htmlquery.Parse(bytes.NewReader(htmlContent))
    if err != nil {
        return nil, fmt.Errorf("parse HTML failed: %v", err)
    }

    // 确定分组节点（支持多层）
    var groups []*html.Node
    if rule.GroupXPath != "" {
        groups, _ = htmlquery.QueryAll(doc, rule.GroupXPath)
    } else {
        groups, _ = htmlquery.QueryAll(doc, rule.RootXPath)
    }

    var cleanedList []model.CleanData

    for _, node := range groups {
        cleaned := model.CleanData{Source: source}

        for fieldName, cfg := range rule.Fields {
            nodeVal := htmlquery.FindOne(node, cfg.XPath)
            if nodeVal == nil {
                continue
            }
            text := htmlquery.InnerText(nodeVal)

            // 如果配置了正则表达式，进行匹配提取
            if cfg.Regex != "" {
                re := regexp.MustCompile(cfg.Regex)
                match := re.FindStringSubmatch(text)
                if len(match) > 1 {
                    text = match[1] // 默认取第一个分组
                }
            }

            // 动态设置字段
            cleaned.Set(fieldName, text)
        }

        cleanedList = append(cleanedList, cleaned)
    }

    return cleanedList, nil
}
