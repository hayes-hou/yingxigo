/*
	Copyright 2019 The Tengyue360 Authors. All rights reserved.
	Ty-Go
	https://jasperxu.github.io/gorm-zh/database.html#dbc
	Model
*/
package model

import (
	"time"
)

// 定义User结构体
type User struct {
	Id          int
	UserName    string
	UserSchool  string
	CreateTime  int
	UpdateTime  int
	DeletedTime *time.Time
}

// 返回UserList
func ListUser() ([]*User, int, error) {
	users := make([]*User, 0)
	var count int
	if err := DB.Master.Model(&User{}).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Master.Order("id DESC").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}
