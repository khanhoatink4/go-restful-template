package main

import (
	"fmt"
	"github.com/khanhoatink4/go-restful-template/models"
	"github.com/gin-gonic/gin"
	"github.com/khanhoatink4/go-restful-template/controllers"
)

func main() {
	db := models.Init()
	defer db.Close()
	fmt.Println("hello")

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		/*** START USER ***/
		article := new(controllers.ArticleController)
		v1.GET("/articles", article.All)
		v1.POST("/article", article.Create)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.Run(":9000")

	/// https://github.com/qiangxue/golang-restful-starter-kit/blob/master/app/init.go
	/// https://github.com/Massad/gin-boilerplate/blob/master/main.go
	/// https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/hello.go
}
