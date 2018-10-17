package models

import "github.com/jinzhu/gorm"

//Article ...
type ArticleModel struct {
	gorm.Model
	ID        int64
	Title     string   `gorm:"column:title"`
	Content   string   `gorm:"column:content"`
}

// set User's table name to be `profiles`
func (ArticleModel) TableName() string {
	return "ArticleModel"
}