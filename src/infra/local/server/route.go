package server

import (
	"github.com/gin-gonic/gin"
)

type (
	Route struct {
		Method     string
		URL        string
		Business   gin.HandlerFunc
		Middleware []gin.HandlerFunc
	}
)
