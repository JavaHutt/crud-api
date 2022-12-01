package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JavaHutt/crud-api/internal/repository"

	"github.com/JavaHutt/crud-api/internal/migrate"
)

func main() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	if err = migrate.Migrate(ctx, db); err != nil {
		log.Fatal(err)
	}
	rep := repository.NewAdvertiseRepo(db)
	fmt.Println(rep.GetAllAdvertise(ctx))
}
