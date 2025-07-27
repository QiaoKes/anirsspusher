package server

var (
    Success = &BaseResponse{Code: "Success", Message: "Success"}
)

type BaseResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

// MediaInfo rss信息结构体
type MediaInfo struct {
    Action            string `json:"action"`            // 执行的动作
    Title             string `json:"title"`             // 中文标题
    JPTitle           string `json:"jpTitle"`           // 日文标题
    Score             string `json:"score"`             // 评分
    TheMovieDBName    string `json:"themoviedbName"`    // TheMovieDB 名称
    TMDBId            string `json:"tmdbid"`            // TMDB ID
    TMDBUrl           string `json:"tmdbUrl"`           // TMDB 链接
    BGMUrl            string `json:"bgmUrl"`            // Bangumi 链接
    Season            string `json:"season"`            // 季号
    Episode           string `json:"episode"`           // 集号
    Subgroup          string `json:"subgroup"`          // 字幕组
    Progress          string `json:"progress"`          // 观看进度
    Premiere          string `json:"premiere"`          // 首播日期
    Text              string `json:"text"`              // 描述文本
    DownloadPath      string `json:"downloadPath"`      // 下载路径
    EpisodeTitle      string `json:"episodeTitle"`      // 剧集标题
    BGMEpisodeTitle   string `json:"bgmEpisodeTitle"`   // Bangumi 剧集标题
    BGMJpEpisodeTitle string `json:"bgmJpEpisodeTitle"` // Bangumi 日文剧集标题
    Image             string `json:"image"`             // 图片链接
}
