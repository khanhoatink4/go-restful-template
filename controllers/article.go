package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khanhoatink4/go-restful-template/models"
)

type ArticleForm struct {
	Title   string `form:"title" json:"title" binding:"required,max=100"`
	Content string `form:"content" json:"content" binding:"required,max=1000"`
}
//ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)

//Create ...
func (ctrl ArticleController) Create(c *gin.Context) {
	var articleForm ArticleForm
	if c.BindJSON(&articleForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": articleForm})
		c.Abort()
		return
	}
	db := models.GetDB()
	err := db.Create(&models.ArticleModel{Title: articleForm.Title, Content: articleForm.Content})
	if err != nil {
		c.JSON(406, gin.H{"message": "Article could not be created", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Article created", "id": err.Value})
}

//All ...
func (ctrl ArticleController) All(c *gin.Context) {
	db := models.GetDB()
	articleModels := []models.ArticleModel{}
	data := db.Find(&articleModels)

	c.JSON(200, gin.H{"data": data.Value})
}