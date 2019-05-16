/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	Handler
*/
package handler

import (
	"github.com/gin-gonic/gin"
	"yingxi.company/infra/go/handler/errno"
	"net/http"
)

// Response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse
func SendResponse(context *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
