package public

import (
	context "backend/src/internal/context/abstract"
	"net/http"
	"time"
)

func Logout(
	ctx context.HandlerContext,
) (interface{}, error) {

	ctx.SetCookie("token", "", time.Now().Add(-1*time.Hour), true, false)

	ctx.Status(http.StatusNoContent)
	return nil, nil
}
