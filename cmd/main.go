package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redis, err := repository.NewRedis(config)
	if err != nil {
		log.Panic(err)
	}
	defer redis.Close()

	ctx := context.Background()
	if err = migrate.Migrate(ctx, db); err != nil {
		log.Panic(err)
	}

	rep := repository.NewQueryRepo(db)
	cache := repository.NewCache(redis, repository.WithExpiration(config.CacheExpiration()))
	querySvc := service.NewQueryService(rep, cache)
	fakerSvc := service.NewFakerService()
	srv := server.NewServer(config, querySvc, fakerSvc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		srv.Shutdown()
	}()

	if err = srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}
}
