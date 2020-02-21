package api

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"

	"pandog/app/api/lib"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	uc "pandog/app/api/usecase/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	DBMS := "mysql"
	DbUser := "root"
	DbPassword := ""
	DbPort := "3306"
	DbHost := "go-db"
	DbName := "go-api"

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal("DB error")
	}

	db.LogMode(true)
	return db
}

func Dispatch() {
	gin.SetMode(gin.DebugMode)
	server := gin.New()

	db := initDB()
	defer db.Close()

	server.GET("/health", func(c *gin.Context) {
		zalo := usecase.Zalo{Db: db}
		// zalo := InitZalo()
		zalo.Taotest()
		c.JSON(http.StatusOK, gin.H{
			"result": "true",
		})
	})

	loadConfig()
	c := lib.NewConfig()
	fmt.Print(c.Get("app.port"))

	port := "4000"
	server.Run(":" + port)
}

func loadConfig() {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename))

	viper.AddConfigPath(filepath + "/config/common")
	viper.SetConfigName("app")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
