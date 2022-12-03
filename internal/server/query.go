//go:generate mockgen -source query.go -destination=./mocks/query.go -package=mocks
package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

const (
	idParam        = "id"
	sortQuery      = "sort"
	pageQuery      = "page"
	statementQuery = "statement"
)

type queryService interface {
	GetAll(ctx context.Context, page int, order string, statement model.QueryStatement) ([]model.SlowestQuery, error)
	Get(ctx context.Context, id int) (*model.SlowestQuery, error)
	Insert(ctx context.Context, query model.SlowestQuery) error
	InsertBulk(ctx context.Context, queries []model.SlowestQuery) error
	Update(ctx context.Context, query model.SlowestQuery) error
	Delete(ctx context.Context, id int) error
}

type queryHandler struct {
	svc queryService
}

func newQueryHandler(svc queryService) queryHandler {
	return queryHandler{
		svc: svc,
	}
}

func (h queryHandler) Routes(router fiber.Router) {
	router.Get("/", h.getAll)
	router.Get("/:id", h.get)
	router.Post("/", h.create)
	router.Put("/:id", h.update)
	router.Delete("/:id", h.delete)
}

// getAll godoc
// @Summary Gets all Query entities
// @Tags    query
// @Produce json
// @Param   sort      query    string false "asc,desc"
// @Param   page      query    int    false "page number, e.g. 2"
// @Param   statement query    string false "select,insert,update,delete"
// @Success 200       {object} []model.SlowestQuery
// @Failure 500
// @Router  /api/v1/query [get]
func (h queryHandler) getAll(c *fiber.Ctx) error {
	sort, err := getSortQuery(c)
	if err != nil {
		return err
	}
	page, err := getPageQuery(c)
	if err != nil {
		return err
	}
	statement, err := getStatementQuery(c)
	if err != nil {
		return err
	}
	res, err := h.svc.GetAll(c.Context(), page, sort, statement)
	if err != nil {
		return encodeError(err)
	}

	return c.JSON(res)
}

// get godoc
// @Summary Get single query entity
// @Tags    query
// @Produce json
// @Param   id  path     int true "id of the ad"
// @Success 200 {object} model.SlowestQuery
// @Failure 500
// @Router  /api/v1/query [get]
func (h queryHandler) get(c *fiber.Ctx) error {
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
// @Summary Creates a single Query entity
// @Tags    query
// @Accept  json
// @Success 201
// @Failure 500
// @Router  /api/v1/query [post]
func (h queryHandler) create(c *fiber.Ctx) error {
	query := new(model.SlowestQuery)
	if err := c.BodyParser(query); err != nil {
		return badRequest(fmt.Sprintf("failed to decode body: %s", err.Error()))
	}

	if err := model.Validate.Struct(query); err != nil {
		return badRequest(err.Error())
	}

	if err := h.svc.Insert(c.Context(), *query); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

// update godoc
// @Summary Update single query entity
// @Tags    query
// @Param   id path int true "id of the query"
// @Success 204
// @Failure 500
// @Router  /api/v1/query [put]
func (h queryHandler) update(c *fiber.Ctx) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	query := new(model.SlowestQuery)
	if err = c.BodyParser(query); err != nil {
		return badRequest(fmt.Sprintf("failed to decode body: %s", err.Error()))
	}

	query.ID = int64(id)
	if err := h.svc.Update(c.Context(), *query); err != nil {
		return encodeError(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// delete godoc
// @Summary Delete single query entity
// @Tags    query
// @Param   id path int true "id of the query"
// @Success 204
// @Failure 500
// @Router  /api/v1/query [delete]
func (h queryHandler) delete(c *fiber.Ctx) error {
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

func getStatementQuery(c *fiber.Ctx) (model.QueryStatement, error) {
	statement := model.QueryStatement(c.Query(statementQuery))
	switch model.QueryStatement(statement) {
	case "",
		model.QueryStatementSelect,
		model.QueryStatementInsert,
		model.QueryStatementUpdate,
		model.QueryStatementDelete:
		return statement, nil
	default:
		return "", badRequest(fmt.Sprintf("bad statement query param: %s", statement))
	}
}
