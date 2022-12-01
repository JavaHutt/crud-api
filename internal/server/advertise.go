package server

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

type advertiseService interface {
	GetAll(ctx context.Context) ([]model.Advertise, error)
}

type advertiseHandler struct {
	svc advertiseService
}

func newAdvertiseHandler(svc advertiseService) advertiseHandler {
	return advertiseHandler{
		svc: svc,
	}
}

func (h advertiseHandler) Routes(router fiber.Router) {
	router.Get("/", h.getAll)
}

func (h advertiseHandler) getAll(ctx *fiber.Ctx) error {
	res, err := h.svc.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}
