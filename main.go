/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	github.com/gin-gonic/gin
	github.com/golang/protobuf/proto
	https://github.com/skyhee/gin-doc-cn
	Main
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"yingxi.company/infra/go/handler/middleware"
	"yingxi.company/infra/go/model"
	"yingxi.company/infra/go/router"
	"yingxi.company/infra/go/util"
)

// Run
func main() {

	// Init Log
	if err := util.InitLog(""); err != nil {
		panic(err)
	}

	// 初始化变量
	var (
		logPath = util.GetKeyByConf("logpath")
		serverPort = util.GetKeyByConf("port")
		use = util.GetKeyByConf("db.use")
	)
	log.InitWithFile(logPath)

	if use == "true" {
		// Init DB
		dbConf := map[string]string{
			"selection": util.GetKeyByConf("db.selection"),
			"master":    util.GetKeyByConf("db.master.host"),
			"slave":     util.GetKeyByConf("db.slave.host"),
		}
		go model.DB.Init(dbConf)
	}

	// Create Gin Engine
	yxWeb := gin.New()

	// Template
	yxWeb.LoadHTMLGlob("template/*")

	// Use Middleware Logger
	yxWeb.Use(middleware.Logger())
	yxWeb.Use(middleware.RequestId())

	// Router
	router.Load(yxWeb)

	// Run
	yxWeb.Run(serverPort)
}
