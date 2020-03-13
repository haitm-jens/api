// +build wireinject

package wire

import (
	hlAuth "pandog/app/api/handler/auth"
	hlUser "pandog/app/api/handler/user"
	ucAuth "pandog/app/api/usecase/auth"
	ucUser "pandog/app/api/usecase/user"
	database "pandog/interface/infra/db"
	repo "pandog/repo"

	"github.com/google/wire"
)

func UserHandlerLoader(db database.MySQL) hlUser.User {
	wire.Build(hlUser.NewUser, ucUser.NewUser, repo.NewRepo)
	return hlUser.User{}
}

func AuthHandlerLoader(db database.MySQL) hlAuth.Auth {
	wire.Build(hlAuth.NewAuth, ucAuth.NewAuth, repo.NewRepo)
	return hlAuth.Auth{}
}
