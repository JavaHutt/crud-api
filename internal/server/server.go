package server

import (
	"github.com/gofiber/fiber/v2"
)

type server struct {
	app      *fiber.App
	port     string
	adSvc    advertiseService
	fakerSvc fakerService
}

func NewServer(name string, port string, adSvc advertiseService, fakerSvc fakerService) server {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: "Fiber",
		// IdleTimeout: 0,
		// ReadTimeout: 0,
		// WriteTimeout: 0,
	})

	return server{
		app:      app,
		port:     port,
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
	v1.Route("/fake", newFakerHandler(s.fakerSvc, s.adSvc).Routes)
}
