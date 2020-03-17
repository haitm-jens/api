package user

import (
	"pandog/domain"
	"pandog/infra/local/server/response"

	"github.com/gin-gonic/gin"
)

type (
	user struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}

	res struct {
		User *user `json:"user"`
	}
)

func toUser(u *domain.User) *user {
	return &user{
		ID:    u.ID,
		Email: u.Email,
	}
}

func toRes(u *domain.User) res {
	t := toUser(u)
	return res{
		User: t,
	}
}

func (s *User) Get(c *gin.Context) {
	r := response.NewJsonResponse(c)

	target := s.uc.GetByID(55)
	data := toRes(target)

	r.Pass(data)
}
