package middleware

import (
	"errors"
	"strings"

	"pandog/app/api/lib/context"
	uc "pandog/app/api/usecase/user"
	"pandog/infra/local/server/response"
	"pandog/repo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		parts := strings.Fields(auth)

		t, err := decodeJwt(parts[1])

		if err != nil {
			abort(c, err)
			return
		}

		uid := uint(t["uid"].(float64))

		check := verifyUser(uid)
		if !check {
			c.Set("AUTH_USER_ID", uid)
			c.Next()
			return
		}

		abort(c, errors.New("ERR_1001"))
	}
}

func abort(c *gin.Context, err error) {
	r := response.NewJsonResponse(c)

	r.Fail(err, nil)
	c.Abort()
}

func decodeJwt(jwtString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ERR_1002")
		}

		c := context.NewConfig()
		salt := []byte(c.Get("auth.salt"))

		return salt, nil
	})

	if str, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return str, nil
	}

	return nil, err
}

func verifyUser(uid uint) bool {
	db := context.NewDb()
	repo := repo.NewRepo(db)
	uc := uc.NewUser(repo)

	target := uc.GetByID(uid)
	return target != nil
}
