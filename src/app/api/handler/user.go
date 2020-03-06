package handler

import (
	"net/http"
	usecase "pandog/app/api/usecase/user"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		Uc usecase.User
	}
)

func NewUser(u usecase.User) User {
	return User{Uc: u}
}

func (u *User) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"me": "me",
	})
}
