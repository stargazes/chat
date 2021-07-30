package models

import (
	"chat/tools"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;unique;comment:主键"`
	CatName   string `gorm:"type:varchar(32);comment:分类名称"`
	CatDesc   string `gorm:"type:varchar(128);comment:分类描述"`
	Sort      int64 `gorm:"comment:排序"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CategoryCreateReq struct {
	CatName string `form:"cat_name" binding:"required"`
	CatDesc string `form:"cat_desc" binding:"required"`
}

var category Category

//同名判断
func FirstCategory(cat_name string) (category Category, err error) {
	if err = tools.Eloquent.Where("cat_name = ?", cat_name).First(&article).Error; err != nil {
		return
	}
	return
}

//创建room
func CreateCategory(category *Category) (id int64, err error) {

	result := tools.Eloquent.Create(category)
	id = category.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
