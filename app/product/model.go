package product

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
)

type ProductModel struct {
	gorm.Model
	Id          string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      string
	AuthorID    uint
	//Tags        []TagModel     `gorm:"many2many:article_tags;"`
	//Comments    []CommentModel `gorm:"ForeignKey:ArticleID"`
}
