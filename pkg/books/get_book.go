package books

import (
	"github.com/andymarthin/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&book)
}