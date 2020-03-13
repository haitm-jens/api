package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"me": "me",
	})
}
