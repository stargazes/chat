package models

import (
	"chat/tools"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID        int64 `gorm:"primaryKey"`
	UserId    int64
	RoomHash  string `gorm:"size:32;unique;index"`
	Name      string `gorm:"size:16;unique"`
	Desc      string `gorm:"type:varchar(32)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CreateReq struct {
	Name string `form:"name" binding:"required"`
	Desc string `form:"desc" binding:"required"`
}

var room Room

//同名判断
func FirstRoom(name string) (room Room, err error) {
	if err = tools.Eloquent.Where("name = ?", name).First(&room).Error; err != nil {
		return
	}
	return
}

//创建room
func CreateRoom(room *Room) (id int64, err error) {

	result := tools.Eloquent.Create(room)
	id = room.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
