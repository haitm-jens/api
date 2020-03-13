package server

import (
	"github.com/gin-gonic/gin"
)

type (
	context struct {
		context *gin.Context
	}
)

func NewContext(c *gin.Context) *context {
	return &context{context: c}
}

func (s *context) Get(key string) interface{} {
	value, flag := s.context.Get(key)
	if flag {
		return value
	}
	return nil
}

func (s *context) FormData(data interface{}) {
	s.context.BindJSON(data)
}
