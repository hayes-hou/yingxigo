/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	go get github.com/spf13/viper
	go get github.com/go-fsnotify/fsnotify
	Util
*/
package util

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

// 配置结构
type Config struct {
	Name string
}

// 初始化
func (conf *Config) getConfig() error {
	// 加载路径
	viper.AddConfigPath("conf")
	viper.SetConfigName("web")
	// 设置格式yaml
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 初始化
func InitLog(cfg string) error {
	conf := Config{
		Name: cfg,
	}
	if err := conf.getConfig(); err != nil {
		return err
	}
	go conf.watchConfig()
	return nil
}

// 热加载配置文件
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("[Web] Config file changed: %s\n", e.Name)
	})
}

// 获取Key值
func GetKeyByConf(key string) string {
	if key == "" {
		return ""
	}
	return viper.GetString(key)
}