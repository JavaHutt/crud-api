package migrations

import (
	"context"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		db.RegisterModel((*model.SlowestQuery)(nil))
		_, err := db.NewCreateTable().Model((*model.SlowestQuery)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_ = db.ResetModel(ctx, (*model.SlowestQuery)(nil))
		_, err := db.NewDropTable().Model((*model.SlowestQuery)(nil)).Exec(ctx)
		return err
	})
}
