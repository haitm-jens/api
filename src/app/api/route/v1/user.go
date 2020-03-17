package v1

import (
	handler "pandog/app/api/handler/user"
	"pandog/app/api/middleware"
	"pandog/infra/local/server"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		handler handler.User
	}
)

func NewUser(h handler.User) User {
	return User{handler: h}
}

func (s User) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "GET",
			URL:      "users/me",
			Business: s.handler.Get,
			Middleware: []gin.HandlerFunc{
				middleware.Auth(),
			},
		},
	}
}
