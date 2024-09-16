package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/milopadma/better-calories-go-backend/internal/days"
	"github.com/milopadma/better-calories-go-backend/internal/fooditems"
	"github.com/milopadma/better-calories-go-backend/internal/mealfooditems"
	"github.com/milopadma/better-calories-go-backend/internal/meals"
	"github.com/milopadma/better-calories-go-backend/internal/users"
)

func New(db *sql.DB) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Handlers
	userHandler := &users.Handler{DB: db}
	dayHandler := &days.Handler{DB: db}
	mealHandler := &meals.Handler{DB: db}
	foodItemHandler := &fooditems.Handler{DB: db}
	mealFoodItemHandler := &mealfooditems.Handler{DB: db}

	// Routes
	users.RegisterRoutes(e, userHandler)
	days.RegisterRoutes(e, dayHandler)
	meals.RegisterRoutes(e, mealHandler)
	fooditems.RegisterRoutes(e, foodItemHandler)
	mealfooditems.RegisterRoutes(e, mealFoodItemHandler)

	return e
}
