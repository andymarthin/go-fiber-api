package books

import (
	"github.com/andymarthin/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBook struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	body := UpdateBookRequestBook{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	return ctx.Status(fiber.StatusOK).JSON(&book)
}