package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	msg struct {
		Err  interface{} `json:"err"`
		Data interface{} `json:"data"`
	}

	err struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	json struct {
		c *gin.Context
	}
)

func NewJsonResponse(context *gin.Context) json {
	return json{c: context}
}

func NewErr(code string, msg string, data interface{}) *err {
	return &err{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func (s *json) Fail(err error, data interface{}, args ...interface{}) {

	status := http.StatusBadRequest
	if len(args) > 0 {
		status = args[0].(int)
	}

	code := err.Error()
	e := NewErr(code, code, data)
	res := msg{
		Err:  e,
		Data: nil,
	}

	s.c.JSON(status, res)
}

func (s *json) Pass(data interface{}, args ...interface{}) {
	res := msg{
		Err:  nil,
		Data: data,
	}

	status := http.StatusOK
	if len(args) > 0 {
		status = args[0].(int)
	}

	s.c.JSON(status, res)
}
