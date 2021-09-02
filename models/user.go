package models

type User struct {
	UID       int64  `gorm:"type:int;primaryKey;not null;autoIncrement"`
	Uuid      string `gorm:"type:varchar(32);not null"`
	Nickname  string `gorm:"type:varchar(32);not null"`
	Headimg   string `gorm:"type:varchar(32)"`
	Status    int8   `gorm:"type:int;not null;default:0"`
	Longitude string `gorm:"type:varchar(32);default:''"`
	Latitude  string `gorm:"type:varchar(32);default:''"`
	CreatedAt int32  `gorm:"autoCreateTime"`
	UpdatedAt int32  `gorm:"autoUpdateTime"`
}
