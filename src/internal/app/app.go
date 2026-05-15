package app

import (
	"log"

	"backend/src/internal/handler/health"
	"backend/src/internal/validator"

	"github.com/gofiber/fiber/v3"
)

func Run() {
	//config := config.Load()
	// conn := postgres.NewPostgresConnection(config.GetDBDSN())
	// serviceProvider := provider.NewServiceProvider()

	//serviceProvider.Register((*sabst.IRosstatService)(nil), simpl.NewRosstatService(conn, rosstatRepo, rosstatAgeRepo))
	//serviceProvider.Register((*sabst.IGeoService)(nil), simpl.NewGeoService(conn, geoRepo))

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

	log.Fatal(app.Listen(":80"))
}
