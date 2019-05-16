/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	https://jasperxu.github.io/gorm-zh/database.html#dbc
	Model
*/
package router

import (
	"github.com/gin-gonic/gin"
	"yingxi.company/infra/go/handler/controller"
)

// Load
func Load(gin *gin.Engine, hf ...gin.HandlerFunc) *gin.Engine {

	// 过滤
	gin.Use(hf...)

	// 控制器
	gin.GET("/home/:name", controller.GetList)
	gin.GET("/test", controller.Test)
	gin.GET("/index", controller.Index)

	return gin
}
