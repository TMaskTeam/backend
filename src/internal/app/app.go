package app

import (
	"log"
	"strconv"

	"backend/src/internal/config"
	"backend/src/internal/db/postgres"
	"backend/src/internal/handler/api"
	"backend/src/internal/handler/health"
	"backend/src/internal/handler/public"
	"backend/src/internal/middleware"
	"backend/src/internal/provider"
	"backend/src/internal/validator"

	sabst "backend/src/internal/service/abstract"
	simpl "backend/src/internal/service/impl"

	rimpl "backend/src/internal/repository/impl"

	"github.com/gofiber/fiber/v3"
)

func Run() {
	config := config.Load()
	conn := postgres.NewPostgresConnection(config.GetDBDSN())
	serviceProvider := provider.NewServiceProvider()

	ownerRepo := rimpl.NewBusinessOwnerRepository()
	serviceProvider.Register((*sabst.IBusinessOwnerService)(nil), simpl.NewBusinessOwnerService(conn, ownerRepo))

	app := fiber.New(fiber.Config{
		EnableSplittingOnParsers: true,
		StructValidator:          validator.NewFiberStructValidator(),
	})

	// CORS middleware
	app.Use(func(c fiber.Ctx) error {
		//c.Set("Access-Control-Allow-Origin", "https://midray.ru")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	app.Get("/ping", health.PingHandler)

	app.Get("/openapi.yaml", api.OpenapiYamlHandler)
	app.Get("/api/*", api.ApiHandler())

	app.Post("/api/v1/auth/owner/register", middleware.Adapt(public.RegisterBusinessOwner, serviceProvider))
	app.Post("/api/v1/auth/client/register", middleware.Adapt(public.RegisterClient, serviceProvider))
	app.Post("/api/v1/auth/owner/register", middleware.Adapt(public.OwnerRegister, serviceProvider))
	app.Post("/api/v1/auth/owner/login", middleware.Adapt(public.OwnerLogin, serviceProvider))

	log.Fatal(app.Listen(":" + strconv.Itoa(config.ServerPort)))
}
