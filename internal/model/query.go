package model

import (
	"time"

	"github.com/uptrace/bun"
)

// QueryStatement is a kind of SQL query
type QueryStatement string

const (
	QueryStatementSelect QueryStatement = "select"
	QueryStatementInsert QueryStatement = "insert"
	QueryStatementUpdate QueryStatement = "update"
	QueryStatementDelete QueryStatement = "delete"
)

// SlowestQuery is a model for SQL query with time spent
type SlowestQuery struct {
	bun.BaseModel `bun:"table:slowest_queries,alias:a"`

	ID        int64          `json:"id" bun:"id,pk,autoincrement"`
	Query     string         `json:"query" bun:",notnull" validate:"required"`
	Statement QueryStatement `json:"statement" bun:",notnull" validate:"required,statementenum"`
	TimeSpent int            `json:"time_spent" bun:",notnull" validate:"required,number"`
	CreatedAt time.Time      `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time      `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty" bun:",soft_delete,nullzero"`
}
