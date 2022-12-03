package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/uptrace/bun"
)

const (
	timeSpentColumn = "time_spent"
	statementColumn = "statement"
	createdAtColumn = "created_at"
	perPage         = 10
)

type queriesRepo struct {
	db *bun.DB
}

// NewQueryRepo is a constructor for queries repository
func NewQueryRepo(db *bun.DB) queriesRepo {
	return queriesRepo{
		db: db,
	}
}

// GetAll selects all the queries
func (rep queriesRepo) GetAll(ctx context.Context, page int, sort string, statement model.QueryStatement) ([]model.SlowestQuery, error) {
	order := "id ASC"
	if sort != "" {
		order = fmt.Sprintf("%s %s, %s", timeSpentColumn, sort, order)
	}

	var queries []model.SlowestQuery
	qb := rep.db.NewSelect().Model(&queries).OrderExpr(order).
		Limit(perPage).Offset((page - 1) * perPage)

	if statement != "" {
		qb.Where("? = ?", bun.Ident(statementColumn), statement)
	}

	if err := qb.Scan(ctx); err != nil {
		return nil, err
	}
	return queries, nil
}

// Get selects single query by it's ID
func (rep queriesRepo) Get(ctx context.Context, id int) (*model.SlowestQuery, error) {
	var query model.SlowestQuery
	if err := rep.db.NewSelect().Model(&query).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &query, nil
}

// Insert creates a single slowest query row
// if the ID is passed and conflict occures, npthing happens
func (rep queriesRepo) Insert(ctx context.Context, query model.SlowestQuery) error {
	_, err := rep.db.NewInsert().Model(&query).Ignore().Exec(ctx)
	return err
}

// InsertBulk creates a multiple query rows
func (rep queriesRepo) InsertBulk(ctx context.Context, queries []model.SlowestQuery) error {
	_, err := rep.db.NewInsert().Model(&queries).Ignore().Exec(ctx)
	return err
}

// Update updates an query by it's ID
func (rep queriesRepo) Update(ctx context.Context, query model.SlowestQuery) error {
	query.UpdatedAt = time.Now().UTC()
	_, err := rep.db.NewUpdate().Model(&query).
		ExcludeColumn(createdAtColumn).
		OmitZero().WherePK().Exec(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNotFound
	}
	return err
}

// Delete deletes an query row by it's ID
func (rep queriesRepo) Delete(ctx context.Context, id int) error {
	if _, err := rep.db.NewDelete().Model((*model.SlowestQuery)(nil)).Where("id = ?", id).Exec(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNotFound
		}
		return err
	}
	return nil
}
