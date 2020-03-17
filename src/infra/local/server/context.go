package server

import (
	"github.com/gin-gonic/gin"
)

type (
	ctx struct {
		c *gin.Context
	}
)

func NewContext(c *gin.Context) *ctx {
	return &ctx{c: c}
}

func (s *ctx) Get(key string) interface{} {
	value, flag := s.c.Get(key)
	if flag {
		return value
	}
	return nil
}

func (s *ctx) FormData(data interface{}) {
	s.c.BindJSON(data)
}
