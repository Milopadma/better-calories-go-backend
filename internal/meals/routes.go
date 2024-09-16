package meals

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.GET("/v1/meals/:dayId", h.GetDayMeals)
	e.POST("/v1/meals", h.CreateMeal)
}