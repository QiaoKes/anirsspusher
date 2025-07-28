package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

// 全局配置结构体
type Config struct {
	Host     string  `toml:"host"`
	Port     int     `toml:"port"`
	Token    string  `toml:"token"`
	Level    int     `toml:"level"`
	GroupIds []int64 `toml:"group_ids"`
	UserIds  []int64 `toml:"user_ids"`
}

var (
	globalConfig *Config
	once         sync.Once
)

// InitConfig 初始化全局配置（线程安全）
func InitConfig(filePath string) error {
	var err error
	once.Do(func() {
		// 解析TOML文件
		var conf Config
		if _, err = toml.DecodeFile(filePath, &conf); err != nil {
			return
		}
		globalConfig = &conf
	})
	return err
}

// GetGlobalConfig 获取全局配置实例（确保先调用InitConfig）
func GetGlobalConfig() *Config {
	return globalConfig
}
