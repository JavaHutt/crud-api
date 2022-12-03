package service

import (
	"context"
	"testing"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/JavaHutt/crud-api/internal/service/mocks"
	"github.com/go-redis/redis/v9"

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
	type repositoryMockData struct {
		page      int
		sort      string
		statement model.QueryStatement

		queries []model.SlowestQuery
		err     error
	}
	testsCases := []struct {
		name           string
		page           int
		sort           string
		statement      model.QueryStatement
		repositoryMock *repositoryMockData

		want []model.SlowestQuery
		err  error
	}{
		{
			name:      "storage error",
			page:      1,
			sort:      "asc",
			statement: model.QueryStatementSelect,
			repositoryMock: &repositoryMockData{
				page:      1,
				sort:      "asc",
				statement: model.QueryStatementSelect,

				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name:      "success",
			page:      1,
			sort:      "asc",
			statement: model.QueryStatementSelect,
			repositoryMock: &repositoryMockData{
				page:      1,
				sort:      "asc",
				statement: model.QueryStatementSelect,

				queries: []model.SlowestQuery{slowQuery, slowestQuery},
			},
			want: []model.SlowestQuery{slowQuery, slowestQuery},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					GetAll(context.Background(), tc.repositoryMock.page, tc.repositoryMock.sort, tc.repositoryMock.statement).
					Return(tc.repositoryMock.queries, tc.repositoryMock.err).
					Times(1)
			}
			svc := NewQueryService(mockRepo, nil)
			actual, err := svc.GetAll(context.Background(), tc.page, tc.sort, tc.statement)

			if tc.err != nil {
				require.Nil(t, actual)
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, tc.want, actual)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type cacheGetMockData struct {
		id string

		query *model.SlowestQuery
		err   error
	}
	type cacheSetMockData struct {
		query *model.SlowestQuery
	}
	type repositoryMockData struct {
		id int

		query *model.SlowestQuery
		err   error
	}
	testsCases := []struct {
		name           string
		id             int
		cacheGetMock   *cacheGetMockData
		cacheSetMock   *cacheSetMockData
		repositoryMock *repositoryMockData

		want *model.SlowestQuery
		err  error
	}{
		{
			name: "query was cached before",
			id:   1,
			cacheGetMock: &cacheGetMockData{
				id:    "1",
				query: &slowQuery,
			},
			want: &slowQuery,
		},
		{
			name: "storage error",
			id:   1,
			cacheGetMock: &cacheGetMockData{
				id: "1",

				err: redis.Nil,
			},
			repositoryMock: &repositoryMockData{
				id: 1,

				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name: "success",
			id:   1,
			cacheGetMock: &cacheGetMockData{
				id: "1",

				err: redis.Nil,
			},
			repositoryMock: &repositoryMockData{
				id:    1,
				query: &slowQuery,
			},
			cacheSetMock: &cacheSetMockData{
				query: &slowQuery,
			},
			want: &slowQuery,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			mockCache := mocks.NewMockcache(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Get(context.Background(), tc.repositoryMock.id).
					Return(tc.repositoryMock.query, tc.repositoryMock.err).
					Times(1)
			}
			if tc.cacheGetMock != nil {
				mockCache.EXPECT().
					Get(context.Background(), tc.cacheGetMock.id).
					Return(tc.cacheGetMock.query, tc.cacheGetMock.err).
					Times(1)
			}
			if tc.cacheSetMock != nil {
				mockCache.EXPECT().
					Set(context.Background(), tc.cacheSetMock.query).
					Return(nil).
					Times(1)
			}
			svc := NewQueryService(mockRepo, mockCache)
			actual, err := svc.Get(context.Background(), tc.id)

			if tc.err != nil {
				require.Nil(t, actual)
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, tc.want, actual)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type repositoryMockData struct {
		query model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name           string
		query          model.SlowestQuery
		repositoryMock *repositoryMockData

		err error
	}{
		{
			name: "storage error",
			repositoryMock: &repositoryMockData{
				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name:  "success",
			query: slowQuery,
			repositoryMock: &repositoryMockData{
				query: slowQuery,
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Insert(context.Background(), tc.repositoryMock.query).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewQueryService(mockRepo, nil)
			err := svc.Insert(context.Background(), tc.query)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInsertBulk(t *testing.T) {
	type repositoryMockData struct {
		queries []model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name           string
		queries        []model.SlowestQuery
		repositoryMock *repositoryMockData

		err error
	}{
		{
			name: "storage error",
			repositoryMock: &repositoryMockData{
				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name:    "success",
			queries: []model.SlowestQuery{slowQuery, slowestQuery},
			repositoryMock: &repositoryMockData{
				queries: []model.SlowestQuery{slowQuery, slowestQuery},
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					InsertBulk(context.Background(), tc.repositoryMock.queries).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewQueryService(mockRepo, nil)
			err := svc.InsertBulk(context.Background(), tc.queries)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type repositoryMockData struct {
		query model.SlowestQuery

		err error
	}
	testsCases := []struct {
		name           string
		query          model.SlowestQuery
		repositoryMock *repositoryMockData

		err error
	}{
		{
			name: "storage error",
			repositoryMock: &repositoryMockData{
				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name:  "success",
			query: slowQuery,
			repositoryMock: &repositoryMockData{
				query: slowQuery,
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Update(context.Background(), tc.repositoryMock.query).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewQueryService(mockRepo, nil)
			err := svc.Update(context.Background(), tc.query)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type repositoryMockData struct {
		id int

		err error
	}
	testsCases := []struct {
		name           string
		id             int
		repositoryMock *repositoryMockData

		err error
	}{
		{
			name: "storage error",
			repositoryMock: &repositoryMockData{
				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name: "success",
			id:   1,
			repositoryMock: &repositoryMockData{
				id: 1,
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockqueryRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Delete(context.Background(), tc.repositoryMock.id).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewQueryService(mockRepo, nil)
			err := svc.Delete(context.Background(), tc.id)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
