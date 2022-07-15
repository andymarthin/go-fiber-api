package books

import (
	"github.com/andymarthin/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBooks(ctx *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&books)
}
