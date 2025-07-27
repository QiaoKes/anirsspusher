package llonebot

const (
    ResponseOk = "ok" // 响应状态：成功
)

const (
    SendGroupMsgApiUri = "/send_group_msg" // 发送群消息API URI
)

const (
    // ErrInvalidRequest 请求无效错误
    ErrInvalidRequest = "invalid request"

    // ErrUnexpectedStatusCode 非预期状态码错误
    ErrUnexpectedStatusCode = "unexpected status code"
)

const (
    Text  = "text"  // 文本消息类型
    Image = "image" // 图片消息类型
)

// SendGroupMsgReq 请求主体结构体
type SendGroupMsgReq struct {
    GroupId int64     `json:"group_id"` // 群组ID
    Message []Message `json:"message"`  // 消息列表
}

// Message 消息结构体
type Message struct {
    Type string      `json:"type"` // 消息类型
    Data MessageData `json:"data"` // 消息数据
}

// MessageData 消息数据结构体
type MessageData struct {
    Text string `json:"text,optional"` // 文本内容
    File string `json:"file,optional"` // 附件内容
}

// SendGroupMsgResp 响应结构体
type SendGroupMsgResp struct {
    Status  string                `json:"status"`  // 响应状态
    RetCode int                   `json:"retcode"` // 响应码
    Data    *SendGroupMsgRespData `json:"data"`    // 响应数据
    Message string                `json:"message"` // 响应消息
    Wording string                `json:"wording"` // 响应措辞
}

type SendGroupMsgRespData struct {
    MessageId int64 `json:"message_id"` // 消息ID
}
