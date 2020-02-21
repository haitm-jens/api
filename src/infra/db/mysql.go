package db

import (
	"fmt"
	"log"

	v "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	MySQL struct {
		conn (*v.DB)
	}
)

func NewMySQL(config map[string]string) *MySQL {
	s := &MySQL{}
	s.Connect(config)
	return s
}

func (s *MySQL) Connect(config map[string]string) {
	DBMS := config["driver"]
	DbUser := config["username"]
	DbPassword := config["password"]
	DbPort := config["port"]
	DbHost := config["host"]
	DbName := config["database"]
	DbCharset := config["charset"]

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName, DbCharset)

	db, err := v.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal("DB error")
	}

	db.LogMode(true)

	s.conn = db
}

func (s *MySQL) Disconnect() {
	s.conn.Close()
}

func (s *MySQL) Find(target interface{}) {
	s.conn.Find(target)
	s.reset()
}

func (s *MySQL) Table(table string) *MySQL {
	s.conn.Table(table)
	return s
}

func (s *MySQL) Limit(limit interface{}) *MySQL {
	s.conn = s.conn.Limit(limit)
	return s
}

func (s *MySQL) Offset(offset interface{}) *MySQL {
	s.conn = s.conn.Offset(offset)
	return s
}

func (s *MySQL) Where(query interface{}, args ...interface{}) *MySQL {
	s.conn = s.conn.Where(query, args...)
	return s
}

func (s *MySQL) Attach(attachments map[string][]interface{}) *MySQL {
	for column, conditions := range attachments {
		s.conn = s.conn.Preload(column, conditions...)
	}
	return s
}

func (s *MySQL) reset() {
	s.conn = s.conn.New()
}
