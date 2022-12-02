package service

import (
	"context"
	"testing"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/JavaHutt/crud-api/internal/service/mocks"

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
	type repositoryMockData struct {
		page int
		sort string

		ads []model.Advertise
		err error
	}
	testsCases := []struct {
		name           string
		page           int
		sort           string
		repositoryMock *repositoryMockData

		want []model.Advertise
		err  error
	}{
		{
			name: "storage error",
			page: 1,
			sort: "asc",
			repositoryMock: &repositoryMockData{
				page: 1,
				sort: "asc",

				err: model.ErrStorage,
			},
			err: model.ErrStorage,
		},
		{
			name: "success",
			page: 1,
			sort: "asc",
			repositoryMock: &repositoryMockData{
				page: 1,
				sort: "asc",

				ads: []model.Advertise{advertise, anotherAdvertise},
			},
			want: []model.Advertise{advertise, anotherAdvertise},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					GetAll(context.Background(), 0, "").
					Return(tc.repositoryMock.ads, tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
			actual, err := svc.GetAll(context.Background(), 0, "")

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
	type repositoryMockData struct {
		id int

		ad  *model.Advertise
		err error
	}
	testsCases := []struct {
		name           string
		id             int
		repositoryMock *repositoryMockData

		want *model.Advertise
		err  error
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
				ad: &advertise,
			},
			want: &advertise,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Get(context.Background(), tc.repositoryMock.id).
					Return(tc.repositoryMock.ad, tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
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
		ad model.Advertise

		err error
	}
	testsCases := []struct {
		name           string
		ad             model.Advertise
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
			ad:   advertise,
			repositoryMock: &repositoryMockData{
				ad: advertise,
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Insert(context.Background(), tc.repositoryMock.ad).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
			err := svc.Insert(context.Background(), tc.ad)

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
		ads []model.Advertise

		err error
	}
	testsCases := []struct {
		name           string
		ads            []model.Advertise
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
			ads:  []model.Advertise{advertise, anotherAdvertise},
			repositoryMock: &repositoryMockData{
				ads: []model.Advertise{advertise, anotherAdvertise},
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					InsertBulk(context.Background(), tc.repositoryMock.ads).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
			err := svc.InsertBulk(context.Background(), tc.ads)

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
		ad model.Advertise

		err error
	}
	testsCases := []struct {
		name           string
		ad             model.Advertise
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
			ad:   advertise,
			repositoryMock: &repositoryMockData{
				ad: advertise,
			},
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Update(context.Background(), tc.repositoryMock.ad).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
			err := svc.Update(context.Background(), tc.ad)

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
			mockRepo := mocks.NewMockadvertiseRepository(ctl)
			if tc.repositoryMock != nil {
				mockRepo.EXPECT().
					Delete(context.Background(), tc.repositoryMock.id).
					Return(tc.repositoryMock.err).
					Times(1)
			}
			svc := NewAdvertiseService(mockRepo)
			err := svc.Delete(context.Background(), tc.id)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
