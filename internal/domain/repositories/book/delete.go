package bookRepository

import (
	"github.com/vitorbgouveia/go-web-api/configs/database"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
)

func DeleteBook(id int) error {
	db := database.GetDatabase()

	var err = db.Delete(&models.Book{}, id).Error
	return err
}