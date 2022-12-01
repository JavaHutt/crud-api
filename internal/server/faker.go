package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

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

func (h fakerHandler) fake(c *fiber.Ctx) error {
	numQuery := c.Query("num", "100")
	num, err := strconv.Atoi(numQuery)
	if err != nil {
		return badRequest(fmt.Sprintf("invalid number query param: %s", numQuery))
	}
	ads := h.svc.Fake(num)
	return h.adSvc.InsertBulk(c.Context(), ads)
}
