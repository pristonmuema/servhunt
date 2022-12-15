package dao

import (
	"time"
)

type User struct {
	ID            int        `gorm:"primary_key; auto_increment" json:"id"`
	FirstName     string     `gorm:"type:varchar(256)" json:"first_name"`
	SecondName    string     `gorm:"type:varchar(256)" json:"second_name"`
	Email         string     `gorm:"type:varchar(256);unique" json:"email"`
	Password      string     `gorm:"not null;unique" json:"password"`
	PhoneNo       string     `gorm:"type:varchar(256);not null;unique" json:"phone_no"`
	Currency      string     `gorm:"type:varchar(256)" json:"currency"`
	Languages     []Language `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"languages"`
	Description   string     `gorm:"type:varchar(256)" json:"description"`
	Ratings       string     `gorm:"type:varchar(256)" json:"ratings"`
	OnlineStatus  string     `gorm:"type:varchar(256)" json:"online_status"`
	UserType      string     `gorm:"type:varchar(256)" json:"user_type"`
	Location      string     `gorm:"type:varchar(256)" json:"location"`
	Address       string     `gorm:"type:varchar(256)" json:"address"`
	AvailableTime string     `gorm:"type:varchar(256)" json:"available_time"`
	About         string     `json:"about"`
	CreatedOn     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	LastUpdatedOn time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"last_updated_on"`
}

type Language struct {
	ID       int    `gorm:"primary_key; auto_increment" json:"id"`
	UserID   int    `json:"user_id"`
	Language string `json:"language"`
}
