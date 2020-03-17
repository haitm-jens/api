package auth

import (
	"time"

	ctx "pandog/app/api/lib/context"
	"pandog/domain"
	"pandog/infra/local/server"
	"pandog/infra/local/server/response"

	libJwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type (
	form struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	res struct {
		Token string `json:"token"`
	}

	jwt struct {
		Id uint `json:"uid"`
		libJwt.StandardClaims
	}
)

func newRes(token string) res {
	return res{
		Token: token,
	}
}

func newJwt(target *domain.Auth) jwt {
	c := ctx.NewConfig()
	lt := c.GetInt("auth.lifetime")

	return jwt{
		Id: target.ID,
		StandardClaims: libJwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, lt, 0).Unix(),
		},
	}
}

func (s *Auth) Login(c *gin.Context) {
	context := server.NewContext(c)

	f := new(form)
	context.FormData(f)

	r := response.NewJsonResponse(c)
	target, err := s.uc.Login(f.Email, f.Password)

	if err != nil {
		r.Fail(err, nil)
		return
	}

	token := generateJwt(target)
	data := newRes(token)

	r.Pass(data)
}

func generateJwt(target *domain.Auth) string {
	c := ctx.NewConfig()
	salt := []byte(c.Get("auth.salt"))
	lib := libJwt.NewWithClaims(libJwt.SigningMethodHS256, newJwt(target))
	token, _ := lib.SignedString(salt)
	return token
}
