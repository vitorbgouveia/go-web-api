package bookRepository

import (
	"github.com/vitorbgouveia/go-web-api/configs/database"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
)

func GetByID(id int) (models.Book, error) {
	db := database.GetDatabase()

	var book models.Book
	var err = db.First(&book, id).Error

	return book, err
}