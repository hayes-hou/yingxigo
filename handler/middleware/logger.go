/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	github.com/lexkong/log
	Middleware
*/
package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"yingxi.company/infra/go/handler"
	"io/ioutil"
	"time"
)

// ResponseWriter Struct
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// ResponseWriter Func
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Start
		runStart := time.Now().UTC()
		urlPath := context.Request.URL.Path

		var bodyBytes []byte
		if context.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(context.Request.Body)
		}

		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		httpMethod := context.Request.Method
		clientIp := context.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: context.Writer,
		}
		context.Writer = blw

		// Continue
		context.Next()

		// Finish
		runFinish := time.Now().UTC()
		runLatency := runFinish.Sub(runStart)

		// Http Return
		var response handler.Response
		code, message := -1, ""
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			// sys
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof("[Web] %s | %s | %s %s | {code: %d, message: %s}", runLatency, clientIp, httpMethod, urlPath, code, message)
	}
}
