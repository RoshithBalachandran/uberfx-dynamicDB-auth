package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roshith/dynamicDB/internals/config"
	"github.com/roshith/dynamicDB/internals/handlers"
	"github.com/roshith/dynamicDB/internals/middleware"
)

func RegisterRoutes(
	app *fiber.App,
	h *handlers.AuthHandlers,
	cfg *config.Config,
) {

	api := app.Group("/api")

	api.Post("/register", h.Register)
	api.Post("/login", h.Login)

	// protected group
	protected := api.Group("/user",
		middleware.JWTMiddleware(cfg.JWT_SECRET),
	)

	protected.Get("/profile", h.Profile)
}
