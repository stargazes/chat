package models

import (
	"chat/tools"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;unique;comment:主键"`
	CatId     int64  `gorm:"comment:分类id"`
	Img       string `gorm:"type:text;comment:头像"`
	Name      string `gorm:"type:varchar(32);comment:名字"`
	Content   string `gorm:"type:text;comment:内容"`
	View      int64  `gorm:"comment:阅读量"`
	Sort      int64  `gorm:"comment:排序"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type ArticleCreateReq struct {
	Name    string `form:"name" binding:"required"`
	Content string `form:"content" binding:"required"`
}
type ArticleListReq struct {
	Page int `form:"p" binding:"required"`
	Size int `form:"r" binding:"required"`
}

var article Article

//同名判断
func FirstArticle(name string) (article Article, err error) {
	if err = tools.Eloquent.Where("name = ?", name).First(&article).Error; err != nil {
		return
	}
	return
}

//创建文章
func CreateArticle(article *Article) (id int64, err error) {

	result := tools.Eloquent.Create(article)
	id = article.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//文章列表
func ArticleList(req *ArticleListReq)(articles Article,err error)  {

	result := tools.Eloquent.Limit(req.Size).Offset(req.Page).Find(&articles)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
