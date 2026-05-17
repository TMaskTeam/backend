package abstract

import "time"

type HandlerContext interface {
	Get(key string, defaultValue ...string) string
	Status(status int) HandlerContext
	BindJSON(data any) error
	SetCookie(name, value string, expiresAt time.Time, httpOnly, secure bool)
	GetLocal(key string) any
}
