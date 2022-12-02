//go:generate mockgen -source advertise.go -destination=./mocks/advertise.go -package=mocks
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
	router.Post("/", h.create)
	router.Put("/:id", h.update)
	router.Delete("/:id", h.delete)
}

func (h advertiseHandler) getAll(c *fiber.Ctx) error {
	res, err := h.svc.GetAll(c.Context())
	if err != nil {
		return encodeError(err)
	}

	return c.JSON(res)
}

func (h advertiseHandler) get(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
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

	if err := h.svc.Insert(c.Context(), *ad); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h advertiseHandler) update(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	ad := new(model.Advertise)
	if err = c.BodyParser(ad); err != nil {
		return badRequest(fmt.Sprintf("failed to decode body: %s", err.Error()))
	}

	ad.ID = int64(id)
	if err := h.svc.Update(c.Context(), *ad); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h advertiseHandler) delete(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	if err = h.svc.Delete(c.Context(), id); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func getIDParam(c *fiber.Ctx) (int, error) {
	idStr := c.Params(idParam)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, badRequest(fmt.Sprintf("invalid id param: %s", idStr))
	}
	return id, nil
}
