package bookRepository

import (
	"github.com/vitorbgouveia/go-web-api/configs/database"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
)

func GetAll() ([]models.Book, error)  {
	db := database.GetDatabase()

	var books []models.Book
	var err = db.Find(&books).Error

	return books, err

}