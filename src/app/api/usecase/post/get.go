package post

import (
	domain "pandog/domain"
)

func (u *Post) Get(id int) {
	domain := new(domain.User)

	u.Db.
		Where("id = ?", id).
		Find(domain)
}
