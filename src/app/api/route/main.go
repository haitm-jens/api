package route

import (
	"strings"

	"pandog/app/api/wire"
	"pandog/interface/infra/db"
	"pandog/interface/route"

	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		db db.MySQL
	}
)

func NewRouter(db db.MySQL) Router {
	return Router{db: db}
}

func (s *Router) Apply(app *gin.Engine) {
	v1 := app.Group("v1")

	collect(v1, wire.AuthRouterLoader(s.db))
	collect(v1, wire.UserRouterLoader(s.db))
}

func collect(v1 *gin.RouterGroup, target route.Base) {
	routes := target.Route()
	for _, rt := range routes {
		switch strings.ToUpper(rt.Method) {
		case "GET":
			v1.
				Use(rt.Middleware...).
				GET(rt.URL, rt.Business)
		case "POST":
			v1.
				Use(rt.Middleware...).
				POST(rt.URL, rt.Business)
		case "PUT":
		case "DELETE":
		default:
		}
	}
}
