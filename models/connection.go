package models

type Connection struct {
	ID         int64 `gorm:"type:int;primaryKey;not null;autoIncrement"`
	UID        int64 `gorm:"type:int;;not null"`
	PartnerUid int64 `gorm:"type:int;not null"`
	Status     int8  `gorm:"type:int;not null;default:0"`
	CreatedAt  int32 `gorm:"autoCreateTime"`
	UpdatedAt  int32 `gorm:"autoUpdateTime"`
}
