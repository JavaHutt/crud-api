//go:generate mockgen -source advertise.go -destination=./mocks/advertise.go -package=mocks
package server

import (
	"bytes"
	"encoding/json"
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

var (
	advertise = model.Advertise{
		ID:       1,
		Name:     "Banner",
		Kind:     model.AdvertiseKindStander,
		Provider: "plista",
		Country:  "Switzerland",
		City:     "Bern",
		Street:   "Main street",
	}
	anotherAdvertise = model.Advertise{
		ID:       2,
		Name:     "Neon",
		Kind:     model.AdvertiseKindNeon,
		Provider: "nativex",
		Country:  "USA",
		City:     "Vice City",
		Street:   "Ocean View",
	}
)

func TestGetAll(t *testing.T) {
	type serviceMockData struct {
		ads []model.Advertise
		err error
	}
	testsCases := []struct {
		name        string
		serviceMock *serviceMockData

		status int
		want   []model.Advertise
	}{
		{
			name: "storage error",
			serviceMock: &serviceMockData{
				err: model.ErrStorage,
			},
			status: http.StatusInternalServerError,
		},
		{
			name: "success",
			serviceMock: &serviceMockData{
				ads: []model.Advertise{advertise, anotherAdvertise},
			},
			status: http.StatusOK,
			want:   []model.Advertise{advertise, anotherAdvertise},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockadvertiseService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					GetAll(gomock.Any()).
					Return(tc.serviceMock.ads, tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newAdvertiseHandler(mockSvc)
			app.Get("/", handler.getAll)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			resp, err := app.Test(req)
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
			if tc.want != nil {
				var ads []model.Advertise
				err = json.NewDecoder(resp.Body).Decode(&ads)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tc.want, ads)
			}
		})
	}
}
func TestGet(t *testing.T) {
	type serviceMockData struct {
		id int

		ad  *model.Advertise
		err error
	}
	testsCases := []struct {
		name        string
		id          int
		serviceMock *serviceMockData

		status int
		want   *model.Advertise
	}{
		{
			name: "storage error",
			id:   1,
			serviceMock: &serviceMockData{
				id:  1,
				err: model.ErrStorage,
			},
			status: http.StatusInternalServerError,
		},
		{
			name: "not found",
			id:   90,
			serviceMock: &serviceMockData{
				id:  90,
				err: model.ErrNotFound,
			},
			status: http.StatusNotFound,
		},
		{
			name: "success",
			id:   1,
			serviceMock: &serviceMockData{
				id: 1,

				ad: &advertise,
			},
			status: http.StatusOK,
			want:   &advertise,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockadvertiseService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Get(gomock.Any(), tc.serviceMock.id).
					Return(tc.serviceMock.ad, tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newAdvertiseHandler(mockSvc)
			app.Get("/:id", handler.get)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%d", tc.id), nil)
			resp, err := app.Test(req)
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
			if tc.want != nil {
				var ad model.Advertise
				err = json.NewDecoder(resp.Body).Decode(&ad)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tc.want, &ad)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type serviceMockData struct {
		ad model.Advertise

		err error
	}
	testsCases := []struct {
		name        string
		ad          model.Advertise
		serviceMock *serviceMockData

		status int
	}{
		{
			name: "storage error",
			serviceMock: &serviceMockData{
				err: model.ErrStorage,
			},
			status: http.StatusInternalServerError,
		},
		{
			name: "success",
			ad:   advertise,
			serviceMock: &serviceMockData{
				ad: advertise,
			},
			status: http.StatusCreated,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockadvertiseService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Insert(gomock.Any(), tc.serviceMock.ad).
					Return(tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newAdvertiseHandler(mockSvc)
			app.Post("/", handler.create)
			body, err := json.Marshal(tc.ad)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req)
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}
