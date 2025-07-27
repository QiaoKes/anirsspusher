package llonebot

import "strconv"

// Options 配置选项
type Options struct {
    host               string // 主机地址
    port               int    // 端口号
    useHttps           bool   // 是否仅使用HTTPS协议
    authorizationToken string // 认证令牌
}

// NewOptions 创建新的配置选项
func NewOptions(host string, port int, token string) *Options {
    return &Options{
        host:               host,
        port:               port,
        authorizationToken: token,
    }
}

// SetHttps 设置为仅使用HTTPS协议
func (o *Options) SetHttps() {
    o.useHttps = true
}

// GetAuthorizationToken 获取认证令牌
func (o *Options) GetAuthorizationToken() string {
    return "Bearer " + o.authorizationToken
}

// GetApiUrl 获取API URI
func (o *Options) GetApiUrl(apiUri string) string {
    protocol := "http"
    if o.useHttps {
        protocol = "https"
    }
    return protocol + "://" + o.host + ":" + strconv.Itoa(o.port) + apiUri
}
