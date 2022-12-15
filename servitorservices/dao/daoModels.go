package dao

import "time"

type Service struct {
	ID              int        `gorm:"primary_key; auto_increment" json:"id"`
	UserID          int        `json:"user_id"`
	ServiceImage    string     `gorm:"type:varchar(256)" json:"service_image"`
	ServiceName     string     `gorm:"type:varchar(256)" json:"service_name"`
	ServiceDuration string     `gorm:"type:varchar(256)" json:"service_duration"`
	ServiceCost     float64    `json:"service_cost"`
	CreatedOn       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	LastUpdatedOn   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"last_updated_on"`
	Category        []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LocationInfo    []Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Location struct {
	ID            int       `gorm:"primary_key; auto_increment" json:"id"`
	ServiceID     int       `json:"service_id"`
	LocationImage string    `gorm:"type:varchar(256)" json:"location_image"`
	LocationName  string    `gorm:"type:varchar(256)" json:"location_name"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Address       string    `gorm:"type:varchar(256)" json:"address"`
	CreatedOn     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	LastUpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_updated_on"`
}

type Category struct {
	ID            int       `gorm:"primary_key; auto_increment" json:"id"`
	ServiceID     int       `json:"service_id"`
	CategoryImage string    `gorm:"type:varchar(256)" json:"category_image"`
	CategoryName  string    `gorm:"type:varchar(256)" json:"category_name"`
	CreatedOn     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_on"`
	LastUpdatedOn time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_updated_on"`
}
