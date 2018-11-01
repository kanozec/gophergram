package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "./../gophergram.db")
	if err != nil {
		fmt.Println("Initialize db connection error", err)
	}
	db.DB().SetMaxIdleConns(10)
	defer db.Close()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":5005")
}
