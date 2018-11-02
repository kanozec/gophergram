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
	c.HTML(http.StatusOK, "type", gin.H{})
}
