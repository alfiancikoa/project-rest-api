package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Price       int    `gorm:"type:int;not null" json:"price" form:"price"`
	Description string `gorm:"type:longtext" json:"description" form:"description"`
	Qty         int    `gorm:"type:int;not null" json:"qty" form:"qty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ListProductRespon struct {
	ID          int
	Name        string
	Price       int
	Description string
	Qty         int
	CreatedAt   string
}
