package server

import (
	"github.com/gofiber/fiber/v2"
)

type server struct {
	app   *fiber.App
	port  string
	adSvc advertiseService
}

func NewServer(name string, port string, adSvc advertiseService) server {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: "Fiber",
		// IdleTimeout: 0,
		// ReadTimeout: 0,
		// WriteTimeout: 0,
	})

	return server{
		app:   app,
		port:  port,
		adSvc: adSvc,
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
	ads := v1.Group("/advertise")
	ads.Route("/", newAdvertiseHandler(s.adSvc).Routes)
}
