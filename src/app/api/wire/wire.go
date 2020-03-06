// +build wireinject

package wire

import (
	"pandog/app/api/handler"

	ucAuth "pandog/app/api/usecase/auth"
	ucUser "pandog/app/api/usecase/user"
	database "pandog/interface/infra/db"
	repo "pandog/repo"

	"github.com/google/wire"
)

func UserHandlerLoader(db database.MySQL) handler.User {
	wire.Build(handler.NewUser, ucUser.NewUser, repo.NewRepo)
	return handler.User{}
}

func AuthHandlerLoader(db database.MySQL) handler.Auth {
	wire.Build(handler.NewAuth, ucAuth.NewAuth, repo.NewRepo)
	return handler.Auth{}
}
