package main

import (
	"context"
	"log"

	"github.com/JavaHutt/crud-api/internal/migrate"
	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/JavaHutt/crud-api/internal/repository"
	"github.com/JavaHutt/crud-api/internal/server"
	"github.com/JavaHutt/crud-api/internal/service"
)

func init() {
	model.RegisterValidators()
}

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
	adSvc := service.NewAdvertiseService(rep)
	fakerSvc := service.NewFakerService()
	srv := server.NewServer("crud-api", ":3000", adSvc, fakerSvc)
	log.Fatal(srv.Start())
}
