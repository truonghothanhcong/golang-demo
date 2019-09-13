package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index:user_username;type:varchar(128);"`
	Email string `gorm:"not null;unique_index:user_email;type:varchar(256);"`
	Password string `gorm:"not null;type:varchar(128);"`
	FullName string `gorm:"type:varchar(256);"`
	//Birthday *time.Time
	IsSuperUser bool `gorm:"default:false;"`
	// AiOi string `sql:"type:varchar(100);not null;unique" json:"name"`
}

type UserProfile struct {
	Id int32 `gorm:"primary_key;auto_increment;"`
	AvatarUrl string `gorm:"size:512;"`
	Address string `gorm:"type:varchar(256);"`
	User User
}
