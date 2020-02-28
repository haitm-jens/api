package user

import (
	domain "pandog/domain"
)

func (u *User) Get(id int) {
	domain := new(domain.User)

	u.Db.
		Where("id = ?", id).
		Find(domain)
}
