package api

import (
	"path"
	"runtime"

	"pandog/app/api/lib/ctx"
	"pandog/app/api/route"
	"pandog/infra/local/lib"

	"github.com/gin-gonic/gin"
)

func Dispatch() {
	c := config()

	switch c.Get("app.mode") {
	case "1":
		gin.SetMode(gin.DebugMode)
	case "0":
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	ctx.NewEnv()
	
	db := ctx.NewDb()
	defer db.Disconnect()

	r := route.NewRouter(db)

	r.Apply(server)

	port := c.Get("app.port")
	server.Run(":" + port)
}

func config() *lib.Config {
	_, filename, _, _ := runtime.Caller(1)
	root := path.Join(path.Dir(filename), "config", "common")

	c := ctx.NewConfig()
	c.Root(root)
	c.LoadDir()

	return c
}
