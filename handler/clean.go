package handler

import (
    "encoding/json"
    "io"
    "net/http"
    "go-cleaning-service/service"
)

// CleanHTMLHandler 处理HTML数据清洗
func CleanHTMLHandler(w http.ResponseWriter, r *http.Request) {
    // 获取源
    source := r.URL.Query().Get("source")
    if source == "" {
        http.Error(w, "missing source parameter", http.StatusBadRequest)
        return
    }

    htmlData, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "failed to read body", http.StatusInternalServerError)
        return
    }

    cleaned, err := service.CleanHTMLData(source, htmlData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cleaned)
}
