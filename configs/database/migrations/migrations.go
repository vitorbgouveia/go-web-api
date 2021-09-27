package migrations

import (
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	"gorm.io/gorm"
)

// RunMigrations execute automatically migrations
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
