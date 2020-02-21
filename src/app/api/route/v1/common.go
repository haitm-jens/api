package v1

import (
	server "pandog/infra/server"
	database "pandog/interface/infra/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	common struct {
	}
)

func NewCommon(db database.MySQL) *common {
	return &common{}
}

func (s *common) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method: "GET",
			URL:    "health",
			Business: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"result": "true",
				})
			},
		},
	}
}
