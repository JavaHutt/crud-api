package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type config interface {
	AppName() string
	APIAddress() string
	IdleTimeout() time.Duration
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
}

type server struct {
	app      *fiber.App
	port     string
	adSvc    advertiseService
	fakerSvc fakerService
}

func NewServer(cfg config, adSvc advertiseService, fakerSvc fakerService) server {
	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName(),
		ServerHeader: "Fiber",
		IdleTimeout:  cfg.IdleTimeout(),
		ReadTimeout:  cfg.ReadTimeout(),
		WriteTimeout: cfg.WriteTimeout(),
	})

	return server{
		app:      app,
		port:     cfg.APIAddress(),
		adSvc:    adSvc,
		fakerSvc: fakerSvc,
	}
}

func (s server) Start() error {
	s.setupRoutes()
	return s.app.Listen(s.port)
}

func (s server) Shutdown() error {
	return s.app.Shutdown()
}

func (s *server) setupRoutes() {
	v1 := s.app.Group("/api/v1")
	v1.Route("/advertise", newAdvertiseHandler(s.adSvc).Routes)
	v1.Route("/faker", newFakerHandler(s.fakerSvc, s.adSvc).Routes)
}
