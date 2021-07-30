package models

import (
	"chat/tools"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         int64  `gorm:"primaryKey;autoIncrement;unique;comment:主键"`
	Pid        int64  `gorm:"comment:父级id"`
	ArticleId  int64  `gorm:"comment:文章id"`
	Message    string `gorm:"type:text;comment:留言内容"`
	UserName   string `gorm:"type:varchar(32);comment:用户名"`
	HeadImg    string `gorm:"type:text;comment:头像地址"`
	ClientIp   string `gorm:"type:varchar(255);comment:客户端ip"`
	ClientInfo string `gorm:"type:text;comment:客户端信息详情"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type MessageCreateReq struct {
	Message string `form:"message" binding:"required"`
}

var comment Comment

//创建room
func CreateComemnt(coment *Comment) (id int64, err error) {

	result := tools.Eloquent.Create(coment)
	id = coment.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
