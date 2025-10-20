package main

import (
	"fmt"
	"go-cleaning-service/handler"
	"go-cleaning-service/service"
	"net/http"
)

func main() {
	// 加载配置
	if err := service.LoadRules("config/rules.json"); err != nil {
		panic(fmt.Sprintf("Failed to load rules: %v", err))
	}

	// 注册接口
	http.HandleFunc("/clean_html", handler.CleanHTMLHandler)

	fmt.Println("🚀 ETL 数据清洗微服务已启动：http://localhost:8080/clean_html")
	http.ListenAndServe(":8080", nil)
}
