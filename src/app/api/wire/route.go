// +build wireinject

package wire

import (
	v1 "pandog/app/api/route/v1"
	database "pandog/interface/infra/db"

	"github.com/google/wire"
)

func UserRouterLoader(db database.MySQL) v1.User {
	wire.Build(v1.NewUser, UserHandlerLoader)
	return v1.User{}
}

func AuthRouterLoader(db database.MySQL) v1.Auth {
	wire.Build(v1.NewAuth, AuthHandlerLoader)
	return v1.Auth{}
}

func PostRouterLoader(db database.MySQL) v1.Post {
	wire.Build(v1.NewPost, PostHandlerLoader)
	return v1.Post{}
}
