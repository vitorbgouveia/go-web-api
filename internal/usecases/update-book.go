package usecases

import (
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	bookRepository "github.com/vitorbgouveia/go-web-api/internal/domain/repositories/book"
)

func UpdateBook(book *models.Book) error {
	var err = bookRepository.UpdateBook(book)

	return err
}