package api

import (
	"log"
	"path"
	"runtime"

	"pandog/app/api/lib"
	route "pandog/app/api/route"
	"pandog/infra/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Dispatch() {
	gin.SetMode(gin.DebugMode)
	server := gin.New()

	dbc := map[string]string{
		"driver":    "mysql",
		"host":      "go-db",
		"port":      "3306",
		"username":  "root",
		"password":  "",
		"database":  "go-api",
		"charset":   "utf8mb4",
		"collation": "utf8mb4_unicode_ci",
	}

	db := db.NewMySQL(dbc)
	defer db.Disconnect()

	r := route.NewRouter(db)

	r.Apply(server)

	loadConfig()
	c := lib.NewConfig()

	port := c.Get("app.port")
	server.Run(":" + port)
}

func loadConfig() {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename))

	// fmt.Println(filepath)
	viper.AddConfigPath(filepath + "/config/common")
	viper.SetConfigName("app")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
