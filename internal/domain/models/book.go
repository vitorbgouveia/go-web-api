package models

import (
	"time"

	"gorm.io/gorm"
)

// Book model book
type Book struct {
	ID          uint           `JSON:id gorm:primaryKey`
	Name        string         `JSON:name`
	Description string         `JSON:description`
	MediumPrice float64        `JSON:mediumPrice`
	Author      string         `JSON:author`
	ImageURL    string         `JSON:imageURL`
	CreatedAt   time.Time      `JSON:createdAt`
	UpdatedAt   time.Time      `JSON:UpdateAt`
	DeletedAt   gorm.DeletedAt `JSON:DeletedAt`
}
