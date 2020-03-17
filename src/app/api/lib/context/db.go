package context

import (
	"sync"

	"pandog/infra/local/db"
)

var db2 *db.MySQL

var once_db2 sync.Once

func NewDb() *db.MySQL {
	once_db2.Do(func() {
		c := NewConfig()
		dbc := c.GetStringMapString("db")

		db2 = db.NewMySQL(dbc)
	})
	return db2
}
