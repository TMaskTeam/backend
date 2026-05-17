package middleware

import (
	"backend/src/pkg/jwt"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func Auth() fiber.Handler {
	return func(fiberCtx fiber.Ctx) error {
		var token string

		token = fiberCtx.Cookies("token")

		if token == "" {
			return fiberCtx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing token"})
		}

		ownerID, role, err := jwt.ValidateToken(token)
		if err != nil {
			return fiberCtx.Status(401).JSON(fiber.Map{"error": "invalid token"})
		}

		fiberCtx.Locals("user_id", ownerID)
		fiberCtx.Locals("role", role)

		return fiberCtx.Next()
	}
}
