package adapter

import (
	"backend/src/internal/context/abstract"
	"time"

	"github.com/gofiber/fiber/v3"
)

type FiberCtxAdapter struct {
	Ctx fiber.Ctx
}

func (fiberAdapter *FiberCtxAdapter) Get(key string, defaultValue ...string) string {
	return fiberAdapter.Ctx.Get(key, defaultValue...)
}

func (fiberAdapter *FiberCtxAdapter) Status(status int) abstract.HandlerContext {
	fiberAdapter.Ctx.Status(status)
	return fiberAdapter
}

func (fiberAdapter *FiberCtxAdapter) BindJSON(data any) error {
	return fiberAdapter.Ctx.Bind().JSON(data)
}

func (fiberAdapter *FiberCtxAdapter) GetLocal(key string) any {
	return fiberAdapter.Ctx.Locals(key)
}

func (fiberAdapter *FiberCtxAdapter) SetCookie(name, value string, expiresAt time.Time, httpOnly, secure bool) {
	fiberAdapter.Ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expiresAt,
		HTTPOnly: httpOnly,
		Secure:   secure,
		SameSite: "Lax",
		Path:     "/",
	})
}
