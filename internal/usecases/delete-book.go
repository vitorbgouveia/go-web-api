package usecases

import bookRepository "github.com/vitorbgouveia/go-web-api/internal/domain/repositories/book"

func DeleteBook(id int) error {
	var err = bookRepository.DeleteBook(id)

	return err
}