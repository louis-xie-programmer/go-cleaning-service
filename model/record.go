package model

type CleanData struct {
    Headline    string `json:"headline,omitempty"`
    BodyText    string `json:"body_text,omitempty"`
    PublishTime string `json:"publish_time,omitempty"`
    Title       string `json:"title,omitempty"`
    Price       string `json:"price,omitempty"`
    URL         string `json:"url,omitempty"`
    Source      string `json:"source,omitempty"`
}

// 动态设置字段值
func (c *CleanData) Set(field string, value string) {
    switch field {
    case "headline":
        c.Headline = value
    case "body_text":
        c.BodyText = value
    case "publish_time":
        c.PublishTime = value
    case "title":
        c.Title = value
    case "price":
        c.Price = value
    case "url":
        c.URL = value
    }
}
