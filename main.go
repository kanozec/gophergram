package main

import (
	"fmt"
	"html/template"
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
		fmt.Println("Initialize db error")
	}
	defer db.Close()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	var templatePage = []string{
		"templates/layouts/frontend.gohtml",
		"templates/layouts/navbar.gohtml",
		"templates/layouts/footer.gohtml",
	}
	templatePage = append(templatePage, "templates/contact.gohtml")
	router.GET("/contact", func(c *gin.Context) {
		router.SetHTMLTemplate(template.Must(template.ParseFiles(templatePage...)))
		c.HTML(200, "frontend", nil)
	})
	router.Run(cfg.ServerPort)
}
