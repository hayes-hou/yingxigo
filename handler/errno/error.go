/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	Error
*/
package errno

import "fmt"

// Error
type Errno struct {
	Code    int
	Message string
}

// Error
func (err Errno) Error() string {
	return err.Message
}

// Err represents an errno
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, errno: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
		case *Err:
			return typed.Code, typed.Message
		case *Errno:
			return typed.Code, typed.Message
		default:
	}

	return InternalServerError.Code, err.Error()
}
