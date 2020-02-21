package route

import (
	v1 "pandog/app/api/route/v1"
	server "pandog/infra/server"
	database "pandog/interface/infra/db"

	"strings"

	"github.com/gin-gonic/gin"
)

type (
	router struct{}
)

func NewRouter() *router {
	return &router{}
}

func (*router) Apply(app *gin.Engine, db database.MySQL) {
	r := collect(db)

	make(app, r)
}

func make(app *gin.Engine, routes []server.Route) {
	v1 := app.Group("v1")
	for _, rt := range routes {
		switch strings.ToUpper(rt.Method) {
		case "GET":
			v1.GET(rt.URL, rt.Business)
		case "POST":
			v1.POST(rt.URL, rt.Business)
		case "PUT":
		case "DELETE":
		default:
		}
	}
}

func collect(db database.MySQL) []server.Route {
	r := []server.Route{}

	r = append(r, (v1.NewCommon(db).Route())...)
	r = append(r, (v1.NewUser(db).Route())...)
	r = append(r, (v1.NewAuth(db).Route())...)
	return r
}
