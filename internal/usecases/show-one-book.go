package usecases

import (
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	bookRepository "github.com/vitorbgouveia/go-web-api/internal/domain/repositories/book"
)

func ShowOneBook(id int) (models.Book, error) {
	var book, err = bookRepository.GetByID(id)

	return book, err
}