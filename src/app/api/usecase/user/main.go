package user

import (
	database "pandog/interface/infra/db"
)

type User struct {
	Db database.MySQL
}

func NewUser(db database.MySQL) User {
	return User{Db: db}
}
