//go:generate mockgen -source faker.go -destination=./mocks/faker.go -package=mocks
package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

const defaultNum = 100

type fakerService interface {
	Fake(num int) []model.Advertise
}

type adsService interface {
	InsertBulk(ctx context.Context, ads []model.Advertise) error
}

type fakerHandler struct {
	svc   fakerService
	adSvc adsService
}

func newFakerHandler(svc fakerService, adSvc adsService) fakerHandler {
	return fakerHandler{
		svc:   svc,
		adSvc: adSvc,
	}
}

func (h fakerHandler) Routes(router fiber.Router) {
	router.Get("/", h.fake)
}

// fake godoc
// @Summary Fake Advertise entities
// @Tags    faker
// @Param   num query int false "number of ads to generate"
// @success 200
// @Router  /fake [get]
func (h fakerHandler) fake(c *fiber.Ctx) error {
	numQuery := c.Query("num", strconv.Itoa(defaultNum))
	num, err := strconv.Atoi(numQuery)
	if err != nil {
		return badRequest(fmt.Sprintf("invalid number query param: %s", numQuery))
	}
	ads := h.svc.Fake(num)
	return h.adSvc.InsertBulk(c.Context(), ads)
}
