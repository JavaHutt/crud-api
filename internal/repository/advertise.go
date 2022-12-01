package repository

import (
	"context"
	"database/sql"
	"errors"

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

// GetAll selects all the advertises
func (rep advertiseRepo) GetAll(ctx context.Context) ([]model.Advertise, error) {
	var ads []model.Advertise
	if err := rep.db.NewSelect().Model(&ads).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return ads, nil
}

// Get selects single ad by it's ID
func (rep advertiseRepo) Get(ctx context.Context, id int) (*model.Advertise, error) {
	var ad model.Advertise
	if err := rep.db.NewSelect().Model(&ad).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &ad, nil
}

// Insert creates a single advertise row
func (rep advertiseRepo) Insert(ctx context.Context, advertise model.Advertise) error {
	_, err := rep.db.NewInsert().Model(advertise).Exec(ctx)
	return err
}

// InsertBulk creates a multiple advertise rows
func (rep advertiseRepo) InsertBulk(ctx context.Context, ads []model.Advertise) error {
	_, err := rep.db.NewInsert().Model(&ads).Exec(ctx)
	return err
}

// Update updates an advertise by it's ID
func (rep advertiseRepo) Update(ctx context.Context, advertise model.Advertise) error {
	_, err := rep.db.NewUpdate().Model(advertise).WherePK().Exec(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNotFound
	}
	return err
}

// Delete deletes an advertise row by it's ID
func (rep advertiseRepo) Delete(ctx context.Context, id int) error {
	_, err := rep.db.NewDelete().Where("id = ?", id).Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNotFound
		}
		return err
	}
	return nil
}
