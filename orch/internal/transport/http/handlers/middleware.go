package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("Auth")
		if err != nil {
			return SendJSON(ctx, ResponseRequiredAuth)
		}

		claims, err := h.services.ParseJWT(cookie.Value)
		if err != nil {
			return SendJSON(ctx, ResponseWrongJWT)
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return SendJSON(ctx, ResponseTokenExpired)
		}

		if claims["uuid"].(string) == "" {
			return SendJSON(ctx, ResponseRequiredAuth)
		}

		ctx.Set("userUUID", claims["uuid"].(string))
		return next(ctx)
	}
}
