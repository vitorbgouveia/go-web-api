package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitorbgouveia/go-web-api/internal/domain/models"
	"github.com/vitorbgouveia/go-web-api/internal/usecases"
)

// ShowBook return one book
func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newID, errID := strconv.Atoi(id)
	if errID != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	var book, err = usecases.ShowOneBook(newID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

// CreateBook add new book
func CreateBook(c *gin.Context) {
	var book models.Book
	c.ShouldBindJSON(&book)

	err := usecases.CreateBook(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

// ShowBooks list all books
func ShowBooks(c *gin.Context) {
	var result, err = usecases.ShowBooks()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})

		return
	}

	c.JSON(200, result)
}

// UpdateBook update one book
func UpdateBook(c *gin.Context) {
	var book models.Book

	errJSON := c.ShouldBindJSON(&book)
	if errJSON != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + errJSON.Error(),
		})

		return
	}

	var err = usecases.UpdateBook(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

// DeleteBook delete one book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newID, errID := strconv.Atoi(id)
	if errID != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	var err = usecases.DeleteBook(newID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
