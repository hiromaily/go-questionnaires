package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//settings
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//url for HTML
	router.GET("/", IndexAction)
	router.GET("/answer", AnswerAction)
	router.GET("/result", ResultAction)

	//url for API


	//Run
	router.Run()
}

// IndexAction is for menu page
func IndexAction(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Index page",
	})
}

// AnswerAction is for answer page
func AnswerAction(c *gin.Context) {
	c.HTML(http.StatusOK, "answer.tmpl", gin.H{
		"title": "Answer Page",
	})
}

// ResultAction is for answer page
func ResultAction(c *gin.Context) {
	c.HTML(http.StatusOK, "result.tmpl", gin.H{
		"title": "Main website",
	})
}

