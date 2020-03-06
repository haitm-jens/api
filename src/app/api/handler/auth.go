package handler

import (
	"net/http"
	usecase "pandog/app/api/usecase/user"

	"github.com/gin-gonic/gin"
)

type (
	Auth struct {
		uc usecase.User
	}
)

func NewAuth(u usecase.User) Auth {
	return Auth{uc: u}
}

func (s *Auth) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "post auth",
	})
}
