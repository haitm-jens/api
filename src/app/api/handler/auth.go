package handler

import (
	"net/http"
	usecase "pandog/app/api/usecase/auth"

	"github.com/gin-gonic/gin"
)

type (
	Auth struct {
		uc usecase.Auth
	}
)

func NewAuth(u usecase.Auth) Auth {
	return Auth{uc: u}
}

func (s *Auth) Login(c *gin.Context) {
	target := s.uc.Login("haitm789@gmail.com", "Aa123456#")

	c.JSON(http.StatusOK, gin.H{
		"result": target,
	})
}
