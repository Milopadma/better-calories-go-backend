package users

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/v1/users", h.CreateUser)
}