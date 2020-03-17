package api

import (
	"fmt"
	"path"
	"runtime"

	lib2 "pandog/app/api/lib"
	"pandog/app/api/route"
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

	db := lib2.NewDb()
	fmt.Printf("h %v\n", db)
	defer db.Disconnect()
	fmt.Printf("h2 %v\n", lib2.NewConfig())
	r := route.NewRouter(db)

	r.Apply(server)

	port := c.Get("app.port")
	server.Run(":" + port)
}

func config() *lib.Config {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename))

	root := filepath + "/config/common"

	c := lib.NewConfig()
	return c.Root(root)
}
