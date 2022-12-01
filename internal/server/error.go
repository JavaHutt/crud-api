package server

import "github.com/gofiber/fiber/v2"

func badRequest(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}
