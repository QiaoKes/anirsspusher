package server

import (
	"anirsspusher/internal/common/config"
	"anirsspusher/internal/common/logger"
	"anirsspusher/pkg/httpClient"
	"anirsspusher/pkg/llonebot"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	llOneBotClient *llonebot.Client
)

func init() {

}

func Start() {
	r := gin.Default()

	conf := config.GetGlobalConfig()
	options := llonebot.NewOptions(conf.Host, conf.Port, conf.Token)
	llOneBotClient = llonebot.NewClient(httpClient.GetHttpClient(), options)

	r.POST("/api/v1/anirss/callback", processAniRssCallback)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func processAniRssCallback(c *gin.Context) {
	// 处理 AniRss 回调逻辑
	// 解析请求体，执行相应的业务逻辑
	requestBody := &MediaInfo{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, BaseResponse{Code: "Error", Message: err.Error()})
		return
	}

	// 遍历发送消息
	groupReqs := buildGroupMessage(config.GetGlobalConfig().GroupIds, requestBody)
	for _, req := range groupReqs {
		_, err := llOneBotClient.SendGroupMessage(req)
		if err != nil {
			logger.Errorf("processAniRssCallback: send group message failed: %v", err)
			c.JSON(500, BaseResponse{Code: "Error", Message: "Failed to send message"})
			return
		}
	}

	PrivateReqs := buildPrivateMessage(config.GetGlobalConfig().UserIds, requestBody)
	for _, req := range PrivateReqs {
		_, err := llOneBotClient.SendPrivateMessage(req)
		if err != nil {
			logger.Errorf("processAniRssCallback: send private message failed: %v", err)
			c.JSON(500, BaseResponse{Code: "Error", Message: "Failed to send message"})
			return
		}
	}

	// 返回响应
	c.JSON(200, Success)
}

func buildGroupMessage(groupIds []int64, info *MediaInfo) []*llonebot.SendGroupMsgReq {
	if info == nil {
		logger.Error("buildMessage: mediaInfo is nil")
		return nil
	}

	msg := buildMessage(info)
	result := make([]*llonebot.SendGroupMsgReq, 0, len(groupIds))
	for i, groupId := range groupIds {
		result[i] = &llonebot.SendGroupMsgReq{
			GroupId: groupId,
			Message: msg,
		}
	}

	return result
}

func buildPrivateMessage(userIds []int64, info *MediaInfo) []*llonebot.SendPrivateMsgReq {
	if info == nil {
		logger.Error("buildMessage: mediaInfo is nil")
		return nil
	}

	msg := buildMessage(info)
	result := make([]*llonebot.SendPrivateMsgReq, 0, len(userIds))
	for i, groupId := range userIds {
		result[i] = &llonebot.SendPrivateMsgReq{
			UserId:  groupId,
			Message: msg,
		}
	}

	return result
}

func buildMessage(info *MediaInfo) []llonebot.Message {
	if info == nil {
		logger.Error("buildMessage: mediaInfo is nil")
		return nil
	}

	return []llonebot.Message{
		llonebot.NewTextMsg("✨ 瞢闇影视更新通知 ✨"),
		llonebot.NewImageMsg(info.Image),
		llonebot.NewTextMsg("🎬 标题: " + info.Title),
		llonebot.NewTextMsg("📺 剧集: " + fmt.Sprintf("S%s-E%s %s", info.Season, info.Episode, info.TheMovieDBName)),
		llonebot.NewTextMsg("⭐ 评分: " + info.Score),
		llonebot.NewTextMsg("🔗 TMDB链接: " + info.TMDBUrl),
		llonebot.NewTextMsg("🔗 BGM链接: " + info.BGMUrl),
		llonebot.NewTextMsg("👥 字幕组: " + info.Subgroup),
		llonebot.NewTextMsg("📊 进度: " + info.Progress),
		llonebot.NewTextMsg("📅 首播: " + info.Premiere),
		llonebot.NewTextMsg("⏱️ 更新时间: " + time.Now().Format("2006-01-02 15:04:05")),
		llonebot.NewTextMsg("🔔 推送类型: " + info.Action),
	}
}
