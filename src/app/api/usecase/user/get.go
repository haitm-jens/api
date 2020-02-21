package user

import (
	domain "pandog/domain"
)

func (u *User) Get() {
	domain := new(domain.User)

	u.Db.
		Where("id = 2").
		Find(domain)
}
