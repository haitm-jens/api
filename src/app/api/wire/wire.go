// +build wireinject

package wire

import (
	handler "pandog/app/api/handler"

	ucUser "pandog/app/api/usecase/user"
	database "pandog/interface/infra/db"

	"github.com/google/wire"
)

func UserHandlerLoader(db database.MySQL) handler.User {
	wire.Build(handler.NewUser, ucUser.NewUser)
	return handler.User{}
}
