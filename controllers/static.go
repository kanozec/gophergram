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

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (s *Static) Contact(c *gin.Context) {
	var articleList = []article{
		article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
		article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	}
	c.HTML(http.StatusOK, "type", gin.H{"payload": articleList})
}
