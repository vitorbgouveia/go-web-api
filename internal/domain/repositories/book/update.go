package bookRepository

import (
	"github.com/vitorbgouveia/go-web-api/configs/database"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
)

func UpdateBook(m *models.Book) error {
	db := database.GetDatabase()

	var err = db.UpdateColumns(&m).Error
	return err
}