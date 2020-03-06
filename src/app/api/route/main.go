package route

import (
	"strings"

	"pandog/app/api/wire"
	server "pandog/infra/local/server"
	database "pandog/interface/infra/db"

	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		db database.MySQL
	}
)

func NewRouter(db database.MySQL) Router {
	return Router{db: db}
}

func (s Router) Apply(app *gin.Engine) {
	r := s.collect()

	make(app, r)
}

func (s Router) collect() []server.Route {
	r := []server.Route{}

	user := wire.UserRouterLoader(s.db)
	auth := wire.AuthRouterLoader(s.db)

	r = append(r, (user.Route())...)
	r = append(r, (auth.Route())...)

	return r
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
