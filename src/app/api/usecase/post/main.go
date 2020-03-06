package post

import (
	database "pandog/interface/infra/db"
)

type Post struct {
	Db database.MySQL
}

func NewPost(db database.MySQL) Post {
	return Post{Db: db}
}
