package main

import (
	"context"
	"log"

	"github.com/JavaHutt/crud-api/config"
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
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redis, err := repository.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer redis.Close()

	ctx := context.Background()
	if err = migrate.Migrate(ctx, db); err != nil {
		log.Fatal(err)
	}

	rep := repository.NewAdvertiseRepo(db)
	cache := repository.NewCache(redis)
	adSvc := service.NewAdvertiseService(rep, cache)
	fakerSvc := service.NewFakerService()
	srv := server.NewServer(config, adSvc, fakerSvc)
	log.Fatal(srv.Start())
}
