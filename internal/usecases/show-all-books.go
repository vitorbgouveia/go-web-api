package usecases

import (
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	bookRepository "github.com/vitorbgouveia/go-web-api/internal/domain/repositories/book"
)

func ShowBooks() ([]models.Book, error) {
	var result, err = bookRepository.GetAll()

	return result, err
}