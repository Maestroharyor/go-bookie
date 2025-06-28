package handlers

import (
	"bookie/data"
	"bookie/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBooksList(c *fiber.Ctx) error {
	return response.Succcess(c, "Success", data.Books)
}

func GetSingleBook(c *fiber.Ctx) error {
	bookIDParam := c.Params("bookId")
	bookID, err := strconv.Atoi(bookIDParam)

	if err != nil {
		return response.Error(c, "Invalid book ID", nil)
	}

	if bookID > len(data.Books) {
		return response.Error(c, "Book not found", nil)
	}

	book := data.Books[bookID]

	return response.Succcess(c, "Success", book)
}

func CreateBook(c *fiber.Ctx) error {
	var body data.Book

	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, "Invalid book data", nil)
	}

	if body.Name == "" || body.Author == "" {
		return response.Error(c, "Invalid book data", nil)
	}

	booksLength := len(data.Books)
	body.ID = booksLength + 1
	data.Books = append(data.Books, body)
	return response.Succcess(c, "Books created successfully", body)
}

func UpdateBook(c *fiber.Ctx) error {
	bookIDParam := c.Params("bookId")
	bookID, err := strconv.Atoi(bookIDParam)

	if err != nil {
		return response.Error(c, "Invalid book ID", nil)
	}

	if bookID > len(data.Books) {
		return response.Error(c, "Book not found", nil)
	}

	var body data.Book

	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, "Invalid book data", nil)
	}

	if body.Name == "" || body.Author == "" {
		return response.Error(c, "Invalid book data", nil)
	}

	body.ID = bookID
	data.Books[bookID] = body

	return response.Succcess(c, "Book updated successfully", body)
}

func DeleteBook(c *fiber.Ctx) error {
	bookIDParam := c.Params("bookId")
	bookID, err := strconv.Atoi(bookIDParam)

	if err != nil {
		return response.Error(c, "Invalid book ID", nil)
	}

	if bookID > len(data.Books) {
		return response.Error(c, "Book not found", nil)
	}

	var body data.Book

	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, "Invalid book data", nil)
	}

	if body.Name == "" || body.Author == "" {
		return response.Error(c, "Invalid book data", nil)
	}

	data.Books = append(data.Books[:bookID], data.Books[bookID+1:]...)
	return response.Succcess(c, "Success", nil)
}
