package server

import (
	"errors"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func encodeError(err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return fiber.NewError(fiber.StatusNotFound)
	default:
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}

func badRequest(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}
