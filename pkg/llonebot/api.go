package llonebot

import (
    "anirsspusher/internal/common/logger"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
)

// Client 与LLOneBot API交互的客户端
type Client struct {
    options    *Options     // 客户端配置选项
    httpClient *http.Client // 可选的HTTP客户端
}

// NewClient 创建一个新的LLOneBot客户端
func NewClient(httpClient *http.Client, options *Options) *Client {
    if options == nil || httpClient == nil {
        panic("httpClient and options must not be nil")
    }

    return &Client{
        httpClient: httpClient,
        options:    options,
    }
}

// SendGroupMessage 发送群消息
func (c *Client) SendGroupMessage(req *SendGroupMsgReq) (*SendGroupMsgRespData, error) {
    if req == nil {
        return nil, errors.New(ErrInvalidRequest)
    }

    // 构建请求URL
    url := c.options.GetApiUrl(SendGroupMsgApiUri)

    body, err := json.Marshal(req)
    if err != nil {
        logger.Errorf("SendGroupMessage: failed to marshal request body: %v", err)
        return nil, err
    }

    // 创建HTTP请求
    httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body)) // 需要添加请求体
    if err != nil {
        return nil, err
    }

    // 设置请求头
    httpReq.Header.Set("Authorization", c.options.GetAuthorizationToken())
    httpReq.Header.Set("Content-Type", "application/json")

    // 发送HTTP请求
    resp, err := c.httpClient.Do(httpReq)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // 检查响应状态码
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(ErrUnexpectedStatusCode)
    }

    // 处理响应体
    result := &SendGroupMsgResp{}
    if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
        return nil, err
    }

    if result.Status != ResponseOk {
        return nil, errors.New(result.Message)
    }

    if result.Data == nil {
        return nil, errors.New("response data is nil")
    }

    return result.Data, nil
}
