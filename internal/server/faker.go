//go:generate mockgen -source faker.go -destination=./mocks/faker.go -package=mocks
package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

const (
	numQuery   = "num"
	defaultNum = 100
)

type fakerService interface {
	Fake(num int) []model.SlowestQuery
}

type querService interface {
	InsertBulk(ctx context.Context, queries []model.SlowestQuery) error
}

type fakerHandler struct {
	svc  fakerService
	qSvc querService
}

func newFakerHandler(svc fakerService, qSvc querService) fakerHandler {
	return fakerHandler{
		svc:  svc,
		qSvc: qSvc,
	}
}

func (h fakerHandler) Routes(router fiber.Router) {
	router.Get("/", h.faker)
}

// faker godoc
// @Summary Fake Query entities
// @Tags    faker
// @Param   num query int false "number of queries to generate"
// @success 200
// @Router  /faker [get]
func (h fakerHandler) faker(c *fiber.Ctx) error {
	numStr := c.Query(numQuery, strconv.Itoa(defaultNum))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return badRequest(fmt.Sprintf("invalid number query param: %s", numStr))
	}
	queries := h.svc.Fake(num)
	return h.qSvc.InsertBulk(c.Context(), queries)
}
