/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	Error
*/
package errno

var (
	OK = &Errno{Code: 0, Message: "OK"}

	InternalServerError = &Errno{Code: 10001, Message: "Internal server errno"}
)
