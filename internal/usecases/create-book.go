package usecases

import (
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	bookRepository "github.com/vitorbgouveia/go-web-api/internal/domain/repositories/book"
)

func CreateBook(book *models.Book) error {
	err := bookRepository.Save(book)

	return err
}