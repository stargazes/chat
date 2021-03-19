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

//添加
func (user User) Insert() (id int64, err error) {
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
