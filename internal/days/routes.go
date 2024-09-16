package days

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/v1/days", h.CreateDay)
	e.GET("/v1/days/:userId", h.GetUserDays)
}