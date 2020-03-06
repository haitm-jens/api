package v1

import (
	"pandog/app/api/handler"
	server "pandog/infra/local/server"
)

type (
	Post struct {
		handler handler.Post
	}
)

func NewPost(h handler.Post) Post {
	return Post{handler: h}
}

func (s *Post) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "GET",
			URL:      "post",
			Business: s.handler.Get,
		},
	}
}
