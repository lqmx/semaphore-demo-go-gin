package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lqmx/log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	err := router.Run()
	if err != nil {
		log.Fatalf("err:%v", err)
	}
}