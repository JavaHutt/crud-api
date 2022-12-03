//go:generate mockgen -source query.go -destination=./mocks/mocks.go -package=mocks
package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JavaHutt/crud-api/internal/model"
)

type queryRepository interface {
	GetAll(ctx context.Context, page int, order string, statement model.QueryStatement) ([]model.SlowestQuery, error)
	Get(ctx context.Context, id int) (*model.SlowestQuery, error)
	Insert(ctx context.Context, query model.SlowestQuery) error
	InsertBulk(ctx context.Context, queries []model.SlowestQuery) error
	Update(ctx context.Context, query model.SlowestQuery) error
	Delete(ctx context.Context, id int) error
}

type cache interface {
	Get(ctx context.Context, id string) (*model.SlowestQuery, error)
	Set(ctx context.Context, query *model.SlowestQuery) error
}

type queryService struct {
	rep   queryRepository
	cache cache
}

// NewQueryService is a constructor for query service
func NewQueryService(rep queryRepository, cache cache) queryService {
	return queryService{
		rep:   rep,
		cache: cache,
	}
}

// GetAll selects all the queries
func (svc queryService) GetAll(ctx context.Context, page int, order string, statement model.QueryStatement) ([]model.SlowestQuery, error) {
	queries, err := svc.rep.GetAll(ctx, page, order, statement)
	if err != nil {
		return nil, fmt.Errorf("failed to get all queries: %w", err)
	}
	return queries, nil
}

// Get selects single query by it's ID
func (svc queryService) Get(ctx context.Context, id int) (*model.SlowestQuery, error) {
	query, err := svc.cache.Get(ctx, strconv.Itoa(id))
	if err == nil {
		return query, nil
	}

	query, err = svc.rep.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get query: %w", err)
	}

	if err = svc.cache.Set(ctx, query); err != nil {
		fmt.Printf("failed to save query to the cache: %s\n", err.Error())
	}

	return query, nil
}

// Insert creates a single query row
func (svc queryService) Insert(ctx context.Context, query model.SlowestQuery) error {
	if err := svc.rep.Insert(ctx, query); err != nil {
		return fmt.Errorf("failed to insert query: %w", err)
	}
	return nil
}

// InsertBulk creates a multiple query rows
func (svc queryService) InsertBulk(ctx context.Context, queries []model.SlowestQuery) error {
	if err := svc.rep.InsertBulk(ctx, queries); err != nil {
		return fmt.Errorf("failed to insert bulk query: %w", err)
	}
	return nil
}

// Update updates an query by it's ID
func (svc queryService) Update(ctx context.Context, query model.SlowestQuery) error {
	if err := svc.rep.Update(ctx, query); err != nil {
		return fmt.Errorf("failed to update query: %w", err)
	}
	return nil
}

// Delete deletes an query row by it's ID
func (svc queryService) Delete(ctx context.Context, id int) error {
	if err := svc.rep.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete query: %w", err)
	}
	return nil
}
