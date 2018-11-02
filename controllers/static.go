package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

func (s *Static) Contact(c *gin.Context) {
	// router.SetHTMLTemplate(template.Must(template.ParseFiles(templatePage...)))
	c.HTML(http.StatusOK, "frontend", gin.H{
		"mail": "LLLL",
	})
}
