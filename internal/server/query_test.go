//go:generate mockgen -source query.go -destination=./mocks/query.go -package=mocks
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/JavaHutt/crud-api/internal/server/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	slowQuery = model.SlowestQuery{
		ID:        1,
		Query:     "SELECT * FROM users",
		Statement: model.QueryStatementSelect,
		TimeSpent: int(5 * time.Second),
	}
	slowestQuery = model.SlowestQuery{
		ID:        2,
		Query:     "UPDATE products SET name = \"scam\" ",
		Statement: model.QueryStatementUpdate,
		TimeSpent: int(10 * time.Second),
	}
)

func TestGetAll(t *testing.T) {
	type serviceMockData struct {
		page int
		sort string

		queries []model.SlowestQuery
		err     error
	}
	testsCases := []struct {
		name        string
		page        string
		sort        string
		serviceMock *serviceMockData

		status int
		want   []model.SlowestQuery
	}{
		{
			name:   "bad sorting query param",
			sort:   "esc",
			status: http.StatusBadRequest,
		},
		{
			name:   "bad page query param",
			sort:   "asc",
			page:   "one",
			status: http.StatusBadRequest,
		},
		{
			name: "storage error",
			page: "1",
			sort: "asc",
			serviceMock: &serviceMockData{
				page: 1,
				sort: "asc",
				err:  model.ErrStorage,
			},
			status: http.StatusInternalServerError,
		},
		{
			name: "success",
			page: "1",
			sort: "asc",
			serviceMock: &serviceMockData{
				page:    1,
				sort:    "asc",
				queries: []model.SlowestQuery{slowQuery, slowestQuery},
			},
			status: http.StatusOK,
			want:   []model.SlowestQuery{slowQuery, slowestQuery},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockqueryService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					GetAll(gomock.Any(), tc.serviceMock.page, tc.serviceMock.sort).
					Return(tc.serviceMock.queries, tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newQueryHandler(mockSvc)
			app.Get("/", handler.getAll)

			target := "/"
			if tc.sort != "" {
				target = fmt.Sprintf("%s?page=%s&sort=%s", target, tc.page, tc.sort)
			}

			req := httptest.NewRequest(http.MethodGet, target, nil)
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
			if tc.want != nil {
				var queries []model.SlowestQuery
				err = json.NewDecoder(resp.Body).Decode(&queries)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tc.want, queries)
			}
		})
	}
}
func TestGet(t *testing.T) {
	type serviceMockData struct {
		id int

		ad  *model.SlowestQuery
		err error
	}
	testsCases := []struct {
		name        string
		id          int
		serviceMock *serviceMockData

		status int
		want   *model.SlowestQuery
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

				ad: &slowQuery,
			},
			status: http.StatusOK,
			want:   &slowQuery,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockqueryService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Get(gomock.Any(), tc.serviceMock.id).
					Return(tc.serviceMock.ad, tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newQueryHandler(mockSvc)
			app.Get("/:id", handler.get)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%d", tc.id), nil)
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
			if tc.want != nil {
				var ad model.SlowestQuery
				err = json.NewDecoder(resp.Body).Decode(&ad)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tc.want, &ad)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	model.RegisterValidators()
	type serviceMockData struct {
		ad model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name        string
		ad          model.SlowestQuery
		serviceMock *serviceMockData

		status int
	}{
		{
			name: "storage error",
			ad:   slowQuery,
			serviceMock: &serviceMockData{
				ad:  slowQuery,
				err: model.ErrStorage,
			},
			status: http.StatusInternalServerError,
		},
		{
			name: "bad query object",
			ad: model.SlowestQuery{
				Query: "SELECT * FROM users",
			},
			status: http.StatusBadRequest,
		},
		{
			name: "success",
			ad:   slowQuery,
			serviceMock: &serviceMockData{
				ad: slowQuery,
			},
			status: http.StatusCreated,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockqueryService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Insert(gomock.Any(), tc.serviceMock.ad).
					Return(tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newQueryHandler(mockSvc)
			app.Post("/", handler.create)
			body, err := json.Marshal(tc.ad)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	type serviceMockData struct {
		ad model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name        string
		id          int
		ad          model.SlowestQuery
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
			name: "not found",
			serviceMock: &serviceMockData{
				err: model.ErrNotFound,
			},
			status: http.StatusNotFound,
		},
		{
			name: "success",
			id:   1,
			ad:   slowQuery,
			serviceMock: &serviceMockData{
				ad: slowQuery,
			},
			status: http.StatusNoContent,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockqueryService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Update(gomock.Any(), tc.serviceMock.ad).
					Return(tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newQueryHandler(mockSvc)
			app.Put("/:id", handler.update)
			body, err := json.Marshal(tc.ad)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%d", tc.id), bytes.NewBuffer(body))
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}

func TestDelete(t *testing.T) {
	type serviceMockData struct {
		id int

		err error
	}
	testsCases := []struct {
		name        string
		id          int
		serviceMock *serviceMockData

		status int
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
			},
			status: http.StatusNoContent,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockSvc := mocks.NewMockqueryService(ctl)

			if tc.serviceMock != nil {
				mockSvc.EXPECT().
					Delete(gomock.Any(), tc.serviceMock.id).
					Return(tc.serviceMock.err).
					Times(1)
			}
			app := fiber.New()
			handler := newQueryHandler(mockSvc)
			app.Delete("/:id", handler.delete)

			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%d", tc.id), nil)
			resp, err := app.Test(req)
			defer func() {
				_ = resp.Body.Close()
			}()
			require.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}
