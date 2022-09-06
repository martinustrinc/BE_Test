package model

import (
	"time"
)

type ProductImage struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Product    Product
	ProductID  string `gorm:"size:36;index"`
	Host       string `gorm:"type:text"`
	Path       string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
