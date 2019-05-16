/*
	Copyright 2019 The Tengyue360 Authors. All rights reserved.
	Ty-Go
	Error
*/
package errno

var (
	OK = &Errno{Code: 0, Message: "OK"}

	InternalServerError = &Errno{Code: 10001, Message: "Internal server errno"}
)
