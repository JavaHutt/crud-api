package service

import (
	"context"
	"fmt"

	"github.com/JavaHutt/crud-api/internal/model"
)

type advertiseRepository interface {
	GetAll(ctx context.Context) ([]model.Advertise, error)
	Get(ctx context.Context, id int) (*model.Advertise, error)
	Insert(ctx context.Context, advertise model.Advertise) error
	InsertBulk(ctx context.Context, ads []model.Advertise) error
	Update(ctx context.Context, advertise model.Advertise) error
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
	ads, err := svc.rep.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all advertise: %w", err)
	}
	return ads, nil

}

func (svc advertiseService) Get(ctx context.Context, id int) (*model.Advertise, error) {
	ad, err := svc.rep.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get advertise: %w", err)
	}
	return ad, nil
}

func (svc advertiseService) Insert(ctx context.Context, ad model.Advertise) error {
	if err := svc.rep.Insert(ctx, ad); err != nil {
		return fmt.Errorf("failed to insert advertise: %w", err)
	}
	return nil
}

func (svc advertiseService) InsertBulk(ctx context.Context, ads []model.Advertise) error {
	if err := svc.rep.InsertBulk(ctx, ads); err != nil {
		return fmt.Errorf("failed to insert bulk advertise: %w", err)
	}
	return nil
}

func (svc advertiseService) Update(ctx context.Context, ad model.Advertise) error {
	if err := svc.rep.Update(ctx, ad); err != nil {
		return fmt.Errorf("failed to update advertise: %w", err)
	}
	return nil
}

func (svc advertiseService) Delete(ctx context.Context, id int) error {
	if err := svc.rep.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete advertise: %w", err)
	}
	return nil
}
