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
	"backend/src/pkg/jwt"

	sabst "backend/src/internal/service/abstract"
	simpl "backend/src/internal/service/impl"

	rimpl "backend/src/internal/repository/impl"

	"github.com/gofiber/fiber/v3"
)

func Run() {
	config := config.Load()
	jwt.InitJWTSecret(config.JWTSecret)

	conn := postgres.NewPostgresConnection(config.GetDBDSN())
	serviceProvider := provider.NewServiceProvider()

	ownerRepo := rimpl.NewBusinessOwnerRepository()
	clientRepo := rimpl.NewClientRepository()
	clientBonusProgramRepo := rimpl.NewClientBonusProgramRepository()
	bonusProgramRepo := rimpl.NewBonusProgramRepository()
	businessRepo := rimpl.NewBusinessRepository()

	serviceProvider.Register((*sabst.IBonusProgramService)(nil), simpl.NewBonusProgramService(conn, bonusProgramRepo))
	serviceProvider.Register((*sabst.IBusinessOwnerService)(nil), simpl.NewBusinessOwnerService(conn, ownerRepo))
	serviceProvider.Register((*sabst.IClientService)(nil), simpl.NewClientService(conn, clientRepo))
	serviceProvider.Register((*sabst.IBusinessService)(nil), simpl.NewBusinessService(conn, businessRepo))
	serviceProvider.Register((*sabst.IClientJoinService)(nil), simpl.NewClientJoinService(conn, clientBonusProgramRepo))

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

	app.Get("/api/v1/businesses/programs", middleware.Adapt(public.GetAllBonusPrograms, serviceProvider))
	app.Get("/api/v1/businesses/:business_id/programs", middleware.Adapt(public.GetBonusProgramsByBusinessID, serviceProvider))
	app.Post("/api/v1/businesses/:business_id/programs", middleware.Adapt(public.CreateBonusProgram, serviceProvider))

	app.Get("/api/*", api.ApiHandler())

	app.Post("/api/v1/auth/owner/register", middleware.Adapt(public.OwnerRegister, serviceProvider))
	app.Post("/api/v1/auth/client/register", middleware.Adapt(public.ClientRegister, serviceProvider))
	app.Post("/api/v1/auth/owner/login", middleware.Adapt(public.OwnerLogin, serviceProvider))
	app.Post("/api/v1/auth/client/login", middleware.Adapt(public.ClientLogin, serviceProvider))

	app.Post("/api/v1/auth/logout", middleware.Auth(), middleware.Adapt(public.Logout, serviceProvider))

	app.Get("/api/v1/me", middleware.Auth(), middleware.Adapt(public.GetMe, serviceProvider))
	app.Put("/api/v1/me", middleware.Auth(), middleware.Adapt(public.Update, serviceProvider))

	app.Post("/api/v1/businesses", middleware.Auth(), middleware.Adapt(public.CreateBusiness, serviceProvider))
	app.Get("/api/v1/businesses", middleware.Auth(), middleware.Adapt(public.GetAllBusinesses, serviceProvider))
	app.Delete("/api/v1/businesses/:business_id", middleware.Auth(), middleware.Adapt(public.DeleteBusiness, serviceProvider))
	app.Post("/api/v1/programs/:program_id/join", middleware.Adapt(public.ClientJoinProgram, serviceProvider))
	app.Get("/api/v1/client/programs", middleware.Adapt(public.GetClientPrograms, serviceProvider))

	app.Get("/api/*", api.ApiHandler())

	log.Fatal(app.Listen(":" + strconv.Itoa(config.ServerPort)))
}
