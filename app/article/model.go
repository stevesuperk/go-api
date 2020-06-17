package article

import (
	"errors"
	"go-api/config"
)

//var db = common.Db

type ArticleModel struct {
	Title    string
	FolderID int64
	MkValue  string
	Tags     string
}

var db = config.GetDBConn()

//Find
func (this ArticleModel) Count() (total int) {
	db.Table("article").Where("deleted=?", 0).Count(&total)
	return total
}

func (this ArticleModel) GetMany(page int) (articles []ArticleModel) {

	db.Table("article").Where("deleted=?", 0).Order("updated_at desc").Offset((page - 1) * 10).Limit(10).
		Select([]string{"id", "updated_at", "title"}).Find(&articles)
	return
}

func (this *ArticleModel) GetArticleInfo() {
	db.Where("id=?", this.ID).First(&this)
}
func (this *ArticleModel) GetArticleInfoByTitle() {
	db.Where("title=?", this.Title).First(&this)
}

func (this ArticleModel) GetDeletedArticle() (articles []ArticleModel) {
	db.Find(&articles, "deleted=?", 1)
	return
}

//Create
func (this *ArticleModel) Add() {
	db.Create(this)
}

//Update Or Create
func (this *ArticleModel) Update() {
	if this.ID != 0 {
		db.Save(&this)
	} else {
		db.Create(this)
	}
}
func (this *ArticleModel) SetTag() {
	db.Model(&this).Update("tags", this.Tags)
}

//回收到垃圾箱
func (this *ArticleModel) Delete() {
	db.Model(&this).Update("deleted", true)
}

//真实批量删除【后台】
func (this ArticleModel) DeleteMany(ids []string) {
	db.Table("article").Where("id in (?)", ids).Delete(&this)
}

//清空垃圾箱
func (this ArticleModel) ClearRubbish() {
	db.Where("deleted <> 0").Delete(&ArticleModel{})
}

//垃圾箱恢复
func (this ArticleModel) Recover() error {
	hasFolder := 0
	db.First(&this)
	db.Table("folder").Where("id=?", this.FolderID).Count(&hasFolder)

	if hasFolder != 0 || this.FolderID == 0 {
		db.Table("article").Where("id=?", this.ID).Update("deleted", 0)
		return nil
	} else {
		return errors.New("父目录不存在！恢复失败")
	}
}

func (this ArticleModel) IsExist() bool {
	c := 0
	db.Table("article").Where("title=?", this.Title).Count(&c)
	return c > 0
}
