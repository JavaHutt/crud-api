package migrations

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		db.RegisterModel((*model.Advertise)(nil))
		_, err := db.NewCreateTable().Model((*model.Advertise)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		db.ResetModel(ctx, (*model.Advertise)(nil))
		_, err := db.NewDropTable().Model((*model.Advertise)(nil)).Exec(ctx)
		return err
	})
}
