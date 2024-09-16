package fooditems

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllFoodItems(c echo.Context) error {
	rows, err := h.DB.Query("SELECT id, name, calories, protein, carbohydrate, fat, sugar FROM food_items")
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching food items"})
	}
	defer rows.Close()

	var foodItems []FoodItem
	for rows.Next() {
		var fi FoodItem
		if err := rows.Scan(&fi.ID, &fi.Name, &fi.Calories, &fi.Protein, &fi.Carbohydrate, &fi.Fat, &fi.Sugar); err != nil {
			fmt.Printf("DEBUG Row scan error: %v\n", err)
			continue
		}
		foodItems = append(foodItems, fi)
	}

	return c.JSON(http.StatusOK, foodItems)
}

func (h *Handler) CreateFoodItem(c echo.Context) error {
	fi := new(FoodItem)
	if err := c.Bind(fi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := h.DB.QueryRow("INSERT INTO food_items (name, calories, protein, carbohydrate, fat, sugar) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		fi.Name, fi.Calories, fi.Protein, fi.Carbohydrate, fi.Fat, fi.Sugar).Scan(&fi.ID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating food item"})
	}

	return c.JSON(http.StatusCreated, fi)
}