package meals

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetDayMeals(c echo.Context) error {
	dayID, err := strconv.ParseInt(c.Param("dayId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid day ID"})
	}

	rows, err := h.DB.Query("SELECT id, name FROM meals WHERE day_id = $1", dayID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching meals"})
	}
	defer rows.Close()

	var meals []Meal
	for rows.Next() {
		var meal Meal
		if err := rows.Scan(&meal.ID, &meal.Name); err != nil {
			fmt.Printf("DEBUG Row scan error: %v\n", err)
			continue
		}
		meal.DayID = dayID
		meals = append(meals, meal)
	}

	return c.JSON(http.StatusOK, meals)
}

func (h *Handler) CreateMeal(c echo.Context) error {
	var meal Meal
	if err := c.Bind(&meal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := h.DB.QueryRow("INSERT INTO meals (day_id, name) VALUES ($1, $2) RETURNING id", meal.DayID, meal.Name).Scan(&meal.ID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating meal"})
	}

	return c.JSON(http.StatusCreated, meal)
}