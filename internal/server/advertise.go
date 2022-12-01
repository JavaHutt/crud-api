package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

const idParam = "id"

type advertiseService interface {
	GetAll(ctx context.Context) ([]model.Advertise, error)
	Get(ctx context.Context, id int) (*model.Advertise, error)
	Insert(ctx context.Context, advertise model.Advertise) error
	InsertBulk(ctx context.Context, ads []model.Advertise) error
	Update(ctx context.Context, advertise model.Advertise) error
	Delete(ctx context.Context, id int) error
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
	router.Get("/:id", h.get)
}

func (h advertiseHandler) getAll(c *fiber.Ctx) error {
	res, err := h.svc.GetAll(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (h advertiseHandler) get(c *fiber.Ctx) error {
	idStr := c.Params(idParam)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return badRequest(fmt.Sprintf("invalid id param: %s", idStr))
	}

	res, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(res)
}
