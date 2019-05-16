/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	github.com/satori/go.uuid
	RequestId
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取Header头
		requestID := context.Request.Header.Get("Ty-Request-Id")
		if requestID == "" {
			uuid4, _ := uuid.NewV4()
			requestID = uuid4.String()
		}
		// Set到Header
		context.Set("RequestId", requestID)
		context.Writer.Header().Set("Ty-Request-Id", requestID)
		context.Next()
	}
}
