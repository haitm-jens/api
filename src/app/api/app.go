package api

import (
	"path"
	"runtime"

	"pandog/app/api/route"
	"pandog/infra/local/db"
	"pandog/infra/local/lib"

	"github.com/gin-gonic/gin"
)

func Dispatch() {
	c := config()
	c.LoadDir()

	switch c.Get("app.mode") {
	case "1":
		gin.SetMode(gin.DebugMode)
	case "0":
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()

	dbc := c.GetStringMapString("db")
	db := db.NewMySQL(dbc)

	defer db.Disconnect()

	r := route.NewRouter(db)

	r.Apply(server)

	port := c.Get("app.port")
	server.Run(":" + port)
}

func config() *lib.Config {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename))

	root := filepath + "/config/common"
	return lib.NewConfig(root)
}
