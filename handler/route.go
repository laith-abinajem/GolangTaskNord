package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) RegisterRoutes() {

	root := s.app.Group("/")

	root.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := root.Group("/api")

	// Business routes
	s.RegisterBusinessAPIs(api)

}

func (s *Server) RegisterBusinessAPIs(router fiber.Router) {
	transactions := router.Group("/transactions")

	transactions.Post("/", s.CreateTransaction)
	transactions.Get("/:tenantID/:productID", s.GetTotalSales)
	transactions.Get("/top-products", s.GetTopProducts)
}
