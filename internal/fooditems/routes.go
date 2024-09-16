package fooditems

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.GET("/v1/food-items", h.GetAllFoodItems)
	e.POST("/v1/food-items", h.CreateFoodItem)
}