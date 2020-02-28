// +build wireinject

package wire

import (
	v1 "pandog/app/api/route/v1"
	database "pandog/interface/infra/db"

	"github.com/google/wire"
)

func UserRoutLoader(db database.MySQL) v1.User {
	wire.Build(v1.NewUser, UserHandlerLoader)
	return v1.User{}
}
