package handler

import (
	"net/http"
	usecase "pandog/app/api/usecase/user"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		uc usecase.User
	}
)

func NewUser(u usecase.User) *User {
	return &User{uc: u}
}

func (u *User) Get(c *gin.Context) {
	u.uc.Get()
	c.JSON(http.StatusOK, gin.H{
		"result": "true",
	})
}

func (u *User) Pandog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "pandog",
	})
}
