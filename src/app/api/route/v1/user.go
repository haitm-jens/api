package v1

import (
	"pandog/app/api/handler"
	usecase "pandog/app/api/usecase/user"
	server "pandog/infra/server"
	database "pandog/interface/infra/db"
)

type (
	user struct {
		handler *handler.User
	}
)

func NewUser(db database.MySQL) *user {
	uc := usecase.NewUser(db)
	hl := handler.NewUser(*uc)
	return &user{handler: hl}
}

func (s *user) Route() []server.Route {
	return []server.Route{
		server.Route{
			Method:   "GET",
			URL:      "users/pandog",
			Business: s.handler.Pandog,
		},
	}
}
