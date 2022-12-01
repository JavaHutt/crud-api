package service

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
)

type advertiseRepository interface {
	GetAll(ctx context.Context) ([]model.Advertise, error)
	Get(ctx context.Context, id int) (*model.Advertise, error)
	Insert(ctx context.Context, advertise *model.Advertise) error
	InsertBulk(ctx context.Context, ads []model.Advertise) error
	Update(ctx context.Context, advertise *model.Advertise) error
	Delete(ctx context.Context, id int) error
}

type advertiseService struct {
	rep advertiseRepository
}

// NewAdvertiseService is a constructor for advertise service
func NewAdvertiseService(rep advertiseRepository) advertiseService {
	return advertiseService{
		rep: rep,
	}
}

func (svc advertiseService) GetAll(ctx context.Context) ([]model.Advertise, error) {
	return svc.rep.GetAll(ctx)
}
