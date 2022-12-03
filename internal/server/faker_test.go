//go:generate mockgen -source faker.go -destination=./mocks/faker.go -package=mocks
package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/JavaHutt/crud-api/internal/server/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFake(t *testing.T) {
	type fakerSvcMockData struct {
		num int

		queries []model.SlowestQuery
	}
	type qSvcMockData struct {
		queries []model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name         string
		num          int
		fakerSvcMock *fakerSvcMockData
		qSvcMock     *qSvcMockData

		status int
	}{
		{
			name: "success",
			num:  2,
			fakerSvcMock: &fakerSvcMockData{
				num:     2,
				queries: []model.SlowestQuery{slowQuery, slowestQuery},
			},
			qSvcMock: &qSvcMockData{
				queries: []model.SlowestQuery{slowQuery, slowestQuery},
			},
			status: http.StatusOK,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			fakerSvc := mocks.NewMockfakerService(ctl)
			qSvc := mocks.NewMockquerService(ctl)

			if tc.fakerSvcMock != nil {
				fakerSvc.EXPECT().
					Fake(tc.fakerSvcMock.num).
					Return(tc.fakerSvcMock.queries).
					Times(1)
			}

			if tc.qSvcMock != nil {
				qSvc.EXPECT().
					InsertBulk(gomock.Any(), tc.qSvcMock.queries).
					Return(tc.qSvcMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newFakerHandler(fakerSvc, qSvc)
			app.Get("/", handler.faker)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?num=%d", tc.num), nil)
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}
