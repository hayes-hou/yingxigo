/*
	Copyright 2019 The Tengyue360 Authors. All rights reserved.
	Ty-Go
	controller User
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"yingxi.company/infra/go/handler"
	"yingxi.company/infra/go/handler/errno"
	"yingxi.company/infra/go/model"
	"net/http"
	"time"
	"math"
	"strings"
)

// 返回结构体
type ListResponse struct {
	UserName   string        `json:"userName"`
	TotalCount int           `json:"totalCount"`
	UserList   []*model.User `json:"userList"`
}

// 获取列表
func GetList(context *gin.Context) {
	userName := context.Param("name")
	users, count, _ := model.ListUser()

	handler.SendResponse(
		context,
		errno.OK,
		ListResponse{
			UserName:   userName,
			TotalCount: count,
			UserList:   users,
		},
	)
}

// Test
func Test(context *gin.Context) {
	context.String(http.StatusOK, "yingxi.company")
}

// 获取当前周
func getWeekDay() []string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	limit := 4 - int(math.Abs(float64(offset)))
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Format("2006-01-02")
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, limit).Format("2006-01-02")
	return []string{weekStart, weekEnd}
}

// Index
func Index(context *gin.Context) {
	weekList := getWeekDay()
	context.HTML(http.StatusOK, "web.html", gin.H{
		"title": "yingxi.company",
		"week": strings.Join(weekList, "---"),
	})
}
