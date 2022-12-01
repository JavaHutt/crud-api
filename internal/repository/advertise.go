package repository

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/uptrace/bun"
)

type advertiseRepo struct {
	db *bun.DB
}

// NewAdvertiseRepo is a constructor for advertise repository
func NewAdvertiseRepo(db *bun.DB) advertiseRepo {
	return advertiseRepo{
		db: db,
	}
}

// GetAllAdvertise selects all the advertises
func (rep advertiseRepo) GetAllAdvertise(ctx context.Context) ([]model.Advertise, error) {
	var ads []model.Advertise
	if err := rep.db.NewSelect().Model(&ads).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return ads, nil
}

// GetAllAdvertise selects single ad by it's ID
func (rep advertiseRepo) GetAdvertise(ctx context.Context, id int) (*model.Advertise, error) {
	var ad model.Advertise
	if err := rep.db.NewSelect().Model(&ad).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return &ad, nil
}

// InsertAdvertise creates a single advertise row
func (rep advertiseRepo) InsertAdvertise(ctx context.Context, advertise *model.Advertise) error {
	_, err := rep.db.NewInsert().Model(advertise).Exec(ctx)
	return err
}

// InsertAdvertiseBulk creates a multiple advertise rows
func (rep advertiseRepo) InsertAdvertiseBulk(ctx context.Context, ads []model.Advertise) error {
	_, err := rep.db.NewInsert().Model(&ads).Exec(ctx)
	return err
}

// UpdateAdvertise updates an advertise by it's ID
func (rep advertiseRepo) UpdateAdvertise(ctx context.Context, advertise *model.Advertise) error {
	_, err := rep.db.NewUpdate().Model(advertise).WherePK().Exec(ctx)
	return err
}

// DeleteAdvertise deletes an advertise row by it's ID
func (rep advertiseRepo) DeleteAdvertise(ctx context.Context, id int) error {
	_, err := rep.db.NewDelete().Where("id = ?", id).Exec(ctx)
	return err
}
