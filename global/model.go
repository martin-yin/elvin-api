package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	CompleteUrl string `json:"complete_url"`
	SimpleUrl   string `json:"simple_url"`
	UserId      int    `json:"user_id"`
	UploadType  int    `json:"upload_type"`
	HappenTime  int    `json:"happen_time"`
	HappenDate  int    `json:"happen_date"`
}
