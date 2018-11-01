package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kanozec/gophergram/model"
)

func main() {
	db, err := model.InitializeDB()
	if err != nil {
		fmt.Println("Initalize db error")
	}
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":5005")
}
