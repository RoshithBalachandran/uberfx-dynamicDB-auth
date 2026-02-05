package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/roshith/dynamicDB/internals/config"
	"github.com/roshith/dynamicDB/internals/database"
	"github.com/roshith/dynamicDB/internals/handlers"
	"github.com/roshith/dynamicDB/internals/repository"
	"github.com/roshith/dynamicDB/internals/routes"
	"github.com/roshith/dynamicDB/internals/service"
	"go.uber.org/fx"
)

func NewFiber() *fiber.App {
	return fiber.New()
}

func StartServer(lc fx.Lifecycle, app *fiber.App, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server running on port:", cfg.APP_PORT)
				if err := app.Listen(":" + cfg.APP_PORT); err != nil {
					log.Println("server error:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("shutting down server...")
			return app.Shutdown()
		},
	})
}

func main() {
	fx.New(

		fx.Provide(
			config.LoadEnv,
			NewFiber,

			database.DatabaseManager,

			repository.NewRepository,

			func(repo repository.UserRepository, cfg *config.Config) *service.AuthService {
				return service.NewAuthService(repo, cfg.JWT_SECRET)
			},

			handlers.NewHandlers,
		),

		fx.Invoke(
			routes.RegisterRoutes,
			StartServer,
		),
	).Run()
}
