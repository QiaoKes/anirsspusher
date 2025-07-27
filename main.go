package main

import (
    "anirsspusher/internal/common/config"
    "anirsspusher/internal/common/logger"
    "anirsspusher/internal/server"
)

func main() {
    if err := config.InitConfig("conf/config.toml"); err != nil {
        logger.Fatalf("init config failed: %v", err)
    }

    // 更改配置
    logger.SetLevel(logger.LogLevel(config.GetGlobalConfig().Level))
    logger.SetColor(true)      // 禁用颜色
    logger.SetCallerInfo(true) // 禁用调用者信息

    // 输出到文件
    if err := logger.SetLogFile("logs/app.log"); err != nil {
        logger.Errorf("create log file failed!: %v", err)
    }

    server.Start()
}
