package bookRepository

import (
	"github.com/vitorbgouveia/go-web-api/configs/database"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
)

func Save(book *models.Book) error  {
	db := database.GetDatabase()

	var err = db.Create(&book).Error
	return err
}