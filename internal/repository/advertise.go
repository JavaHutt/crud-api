package repository

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/uptrace/bun"
)

type advertiseRepo struct {
	db *bun.DB
}

func NewAdvertiseRepo(db *bun.DB) advertiseRepo {
	return advertiseRepo{
		db: db,
	}
}

func (rep advertiseRepo) GetAllAdvertise(ctx context.Context) ([]model.Advertise, error) {
	var ads []model.Advertise
	if err := rep.db.NewSelect().Model(&ads).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return ads, nil
}
