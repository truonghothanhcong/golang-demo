package models

import "github.com/jinzhu/gorm"

type Discussion struct {
	gorm.Model
	Name string `gorm:"type:varchar(256);not null;"`
	Content string `gorm:"not null;"`
	UserCreated User `gorm:"not null;"`
	Views int32 `gorm:"default:0;"`
	Votes int32 `gorm:"default:0;"`
	Reports int32 `gorm:"default:0;"`
	IsDeleted bool `gorm:"default:false;"`
}
