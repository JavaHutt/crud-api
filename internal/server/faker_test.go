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

		ads []model.Advertise
	}
	type adsSvcMockData struct {
		ads []model.Advertise

		err error
	}
	testsCases := []struct {
		name         string
		num          int
		fakerSvcMock *fakerSvcMockData
		adsSvcMock   *adsSvcMockData

		status int
	}{
		{
			name: "success",
			num:  2,
			fakerSvcMock: &fakerSvcMockData{
				num: 2,
				ads: []model.Advertise{advertise, anotherAdvertise},
			},
			adsSvcMock: &adsSvcMockData{
				ads: []model.Advertise{advertise, anotherAdvertise},
			},
			status: http.StatusOK,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			fakerSvc := mocks.NewMockfakerService(ctl)
			adsSvc := mocks.NewMockadsService(ctl)

			if tc.fakerSvcMock != nil {
				fakerSvc.EXPECT().
					Fake(tc.fakerSvcMock.num).
					Return(tc.fakerSvcMock.ads).
					Times(1)
			}

			if tc.adsSvcMock != nil {
				adsSvc.EXPECT().
					InsertBulk(gomock.Any(), tc.adsSvcMock.ads).
					Return(tc.adsSvcMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newFakerHandler(fakerSvc, adsSvc)
			app.Get("/", handler.fake)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?num=%d", tc.num), nil)
			resp, err := app.Test(req)
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}
