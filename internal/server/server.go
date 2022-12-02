package server

import (
	"time"

	_ "github.com/JavaHutt/crud-api/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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

// @title       Fiber CRUD Example API
// @version     1.0
// @author      Alexander Karle
// @description This is a sample server server.

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3000
// @BasePath /
// @schemes  http
func NewServer(cfg config, adSvc advertiseService, fakerSvc fakerService) server {
	app := fiber.New(fiber.Config{
		AppName:        cfg.AppName(),
		ServerHeader:   "Fiber",
		IdleTimeout:    cfg.IdleTimeout(),
		ReadTimeout:    cfg.ReadTimeout(),
		WriteTimeout:   cfg.WriteTimeout(),
		ReadBufferSize: 8192,
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
	s.app.Get("/swagger/*", swagger.HandlerDefault)
	v1 := s.app.Group("/api/v1")
	v1.Route("/advertise", newAdvertiseHandler(s.adSvc).Routes)
	v1.Route("/faker", newFakerHandler(s.fakerSvc, s.adSvc).Routes)
}
