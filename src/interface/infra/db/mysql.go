package db

import (
	"pandog/infra/local/db"
)

type MySQL interface {
	Connect(map[string]string)
	Disconnect()
	Limit(limit interface{}) *db.MySQL
	Offset(offset interface{}) *db.MySQL
	Where(query interface{}, args ...interface{}) *db.MySQL
	Attach(attachments map[string][]interface{}) *db.MySQL
	Find(target interface{})
}
