// +build wireinject

package wire

import (
	handler "pandog/app/api/handler"

	ucPost "pandog/app/api/usecase/post"
	ucUser "pandog/app/api/usecase/user"
	database "pandog/interface/infra/db"

	"github.com/google/wire"
)

func UserHandlerLoader(db database.MySQL) handler.User {
	wire.Build(handler.NewUser, ucUser.NewUser)
	return handler.User{}
}

func PostHandlerLoader(db database.MySQL) handler.Post {
	wire.Build(handler.NewPost, ucPost.NewPost)
	return handler.Post{}
}

func AuthHandlerLoader(db database.MySQL) handler.Auth {
	wire.Build(handler.NewAuth, ucUser.NewUser)
	return handler.Auth{}
}
