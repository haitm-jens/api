package repo

import (
	"pandog/interface/domain"
	"pandog/interface/infra/db"
)

type (
	Repo struct {
		db (db.MySQL)
	}
)

func NewRepo(db db.MySQL) Repo {
	return Repo{db: db}
}

func (r *Repo) GetByID(target domain.Base, id uint) {
	r.
		Limit(1).
		Where("id = ?", id).
		Find(target)
}

func (r *Repo) GetByFieldName(target domain.Base, field string, value string) {
	r.
		Limit(1).
		Where(field+" = ?", value).
		Find(target)
}

func (r *Repo) Find(target interface{}) {
	r.db.Find(target)
}

func (r *Repo) Where(query interface{}, args ...interface{}) *Repo {
	r.db = r.db.Where(query, args...)
	return r
}

func (r *Repo) Limit(limit int) *Repo {
	r.db = r.db.Limit(limit)
	return r
}

func (r *Repo) Offset(offset int) *Repo {
	r.db = r.db.Offset(offset)
	return r
}

func (r *Repo) Attach(attachments map[string][]interface{}) *Repo {
	r.db = r.db.Attach(attachments)

	return r
}
