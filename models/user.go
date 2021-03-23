package models

import (
	"chat/tools"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"size:16"`
	Password  string `gorm:"type:varchar(32)"`
	Phone     string `gorm:"type:varchar(11)`
	Token     string `gorm:"type:varchar(32)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

var Users []User

type LoginReq struct {
	Phone string `form:"phone" binding:"required"`
	Pwd   string `form:"pwd"" binding:"required"`
}

type RegisterReq struct {
	Phone    string `form:"phone" binding:"required"`
	Pwd      string `form:"pwd" binding:"required"`
	UserName string `form:"user_name" binding:"required"`
}

/**
登录检查
*/
func LoginCheck(req LoginReq) (bool, User, error, string) {

	var user User

	err := tools.Eloquent.Where("phone = ?", req.Phone).First(&user).Error

	if err == nil {

		if user.Password == req.Pwd {

			return true, user, err, ""

		} else {

			return false, user, err, "密码错误"
		}

	} else {

		return false, user, err, "账号错误"
	}
}

//添加
func Insert(user User) (id int64, err error) {
	result := tools.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = tools.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}
