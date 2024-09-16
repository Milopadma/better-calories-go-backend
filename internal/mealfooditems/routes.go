package mealfooditems

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.GET("/v1/meal-food-items/:mealId", h.GetMealFoodItems)
	e.POST("/v1/meal-food-items", h.CreateMealFoodItem)
}