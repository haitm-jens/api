package handler

import (
	"net/http"
	usecase "pandog/app/api/usecase/post"

	"github.com/gin-gonic/gin"
)

type (
	Post struct {
		Uc usecase.Post
	}
)

func NewPost(u usecase.Post) Post {
	return Post{Uc: u}
}

func (s *Post) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "post auth",
	})
}
