//go:generate mockgen -source advertise.go -destination=./mocks/advertise.go -package=mocks
package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

const (
	idParam   = "id"
	sortQuery = "sort"
	pageQuery = "page"
)

type advertiseService interface {
	GetAll(ctx context.Context, page int, order string) ([]model.Advertise, error)
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

// getAll godoc
// @Summary Gets all Advertise entities
// @Tags    advertise
// @Produce json
// @Param   sort query    string false "asc,desc"
// @Param   page query    int    false "page number, e.g. 2"
// @Success 200  {object} []model.Advertise
// @Failure 500
// @Router  /api/v1/advertise [get]
func (h advertiseHandler) getAll(c *fiber.Ctx) error {
	sort, err := getSortQuery(c)
	if err != nil {
		return err
	}
	page, err := getPageQuery(c)
	if err != nil {
		return err
	}
	res, err := h.svc.GetAll(c.Context(), page, sort)
	if err != nil {
		return encodeError(err)
	}

	return c.JSON(res)
}

// get godoc
// @Summary Get single Advertise entity
// @Tags    advertise
// @Produce json
// @Param   id  path     int true "id of the ad"
// @Success 200 {object} model.Advertise
// @Failure 500
// @Router  /api/v1/advertise [get]
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

// create godoc
// @Summary Creates a single Advertise entity
// @Tags    advertise
// @Accept  json
// @Success 201
// @Failure 500
// @Router  /api/v1/advertise [post]
func (h advertiseHandler) create(c *fiber.Ctx) error {
	ad := new(model.Advertise)
	if err := c.BodyParser(ad); err != nil {
		return badRequest(fmt.Sprintf("failed to decode body: %s", err.Error()))
	}

	if err := model.Validate.Struct(ad); err != nil {
		return badRequest(err.Error())
	}

	if err := h.svc.Insert(c.Context(), *ad); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

// update godoc
// @Summary Update single Advertise entity
// @Tags    advertise
// @Param   id path int true "id of the ad"
// @Success 204
// @Failure 500
// @Router  /api/v1/advertise [put]
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

// delete godoc
// @Summary Delete single Advertise entity
// @Tags    advertise
// @Param   id path int true "id of the ad"
// @Success 204
// @Failure 500
// @Router  /api/v1/advertise [delete]
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

func getSortQuery(c *fiber.Ctx) (string, error) {
	sort := c.Query(sortQuery)
	if sort == "" {
		return "", nil
	}
	if sort != "asc" && sort != "desc" {
		return "", badRequest(fmt.Sprintf("bad sort query param: %s", sort))
	}
	return sort, nil
}

func getPageQuery(c *fiber.Ctx) (int, error) {
	pageStr := c.Query(pageQuery, "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, badRequest(fmt.Sprintf("invalid page param: %s", pageStr))
	}
	return page, nil
}
