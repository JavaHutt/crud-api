package server

import (
	"context"
	"fmt"
	"log"
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
	router.Post("/", h.create)
}

func (h advertiseHandler) getAll(c *fiber.Ctx) error {
	res, err := h.svc.GetAll(c.Context())
	if err != nil {
		return encodeError(err)
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
		return encodeError(err)
	}

	return c.JSON(res)
}

func (h advertiseHandler) create(c *fiber.Ctx) error {
	ad := new(model.Advertise)
	if err := c.BodyParser(ad); err != nil {
		return badRequest(fmt.Sprintf("failed to decode body: %s", err.Error()))
	}
	log.Fatal(ad)
	if err := h.svc.Insert(c.Context(), *ad); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}
