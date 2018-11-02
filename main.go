package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kanozec/gophergram/model"
)

type config struct {
	ServerPort string `env:"SERVERPORT" envDefault:":8080"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env not found, using the default value")
	}

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println("Failed to parse ENV")
	}

	db, err := model.InitializeDB()
	if err != nil {
		fmt.Println("Initalize db error")
	}
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(cfg.ServerPort)
}
