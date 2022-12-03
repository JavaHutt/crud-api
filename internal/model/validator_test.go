package model

import (
	"testing"
	"time"
)

const (
	query = "SELECT * FROM users"
	spent = int(5 * time.Second)
)

func TestValidateSlowestQuery(t *testing.T) {
	RegisterValidators()
	type fields struct {
		ID        int64
		Query     string
		Statement QueryStatement
		Timespent int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "no query",
			fields: fields{
				Statement: QueryStatementInsert,
				Timespent: int(5 * time.Second),
			},
			wantErr: true,
		},
		{
			name: "no statement",
			fields: fields{
				Query:     query,
				Timespent: spent,
			},
			wantErr: true,
		},
		{
			name: "unknown statement",
			fields: fields{
				Query:     query,
				Statement: "upsert",
				Timespent: spent,
			},
			wantErr: true,
		},
		{
			name: "no time spent",
			fields: fields{
				Query:     query,
				Statement: QueryStatementDelete,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Query:     query,
				Statement: QueryStatementDelete,
				Timespent: spent,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := SlowestQuery{
				ID:        tc.fields.ID,
				Query:     tc.fields.Query,
				Statement: tc.fields.Statement,
				TimeSpent: tc.fields.Timespent,
			}
			if err := Validate.Struct(a); (err != nil) != tc.wantErr {
				t.Errorf("SlowestQuery.Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
