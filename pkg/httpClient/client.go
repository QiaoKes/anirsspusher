package httpClient

import (
    "net/http"
    "sync"
)

var (
    httpClient *http.Client // 全局HTTP客户端实例
    once       sync.Once    // 确保只初始化一次
)

// GetHttpClient 获取全局HTTP客户端实例
func GetHttpClient() *http.Client {
    once.Do(func() {
        httpClient = &http.Client{
            Transport: &http.Transport{
                // 可以在这里设置其他HTTP客户端选项
            },
        }
    })

    return httpClient
}
