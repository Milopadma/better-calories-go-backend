package mealfooditems

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

func (h *Handler) GetMealFoodItems(c echo.Context) error {
	mealID, err := strconv.ParseInt(c.Param("mealId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meal ID"})
	}

	rows, err := h.DB.Query(`
		SELECT mfi.id, mfi.quantity, fi.id, fi.name, fi.calories, fi.protein, fi.carbohydrate, fi.fat, fi.sugar
		FROM meal_food_items mfi
		INNER JOIN food_items fi ON mfi.food_item_id = fi.id
		WHERE mfi.meal_id = $1`, mealID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching meal food items"})
	}
	defer rows.Close()

	var mealFoodItems []MealFoodItem
	for rows.Next() {
		var mfi MealFoodItem
		if err := rows.Scan(&mfi.ID, &mfi.Quantity, &mfi.FoodItem.ID, &mfi.FoodItem.Name, &mfi.FoodItem.Calories, &mfi.FoodItem.Protein, &mfi.FoodItem.Carbohydrate, &mfi.FoodItem.Fat, &mfi.FoodItem.Sugar); err != nil {
			fmt.Printf("DEBUG Row scan error: %v\n", err)
			continue
		}
		mfi.MealID = mealID
		mealFoodItems = append(mealFoodItems, mfi)
	}

	return c.JSON(http.StatusOK, mealFoodItems)
}

func (h *Handler) CreateMealFoodItem(c echo.Context) error {
	mfi := new(MealFoodItem)
	if err := c.Bind(mfi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := h.DB.QueryRow("INSERT INTO meal_food_items (meal_id, food_item_id, quantity) VALUES ($1, $2, $3) RETURNING id",
		mfi.MealID, mfi.FoodItem.ID, mfi.Quantity).Scan(&mfi.ID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating meal food item"})
	}

	return c.JSON(http.StatusCreated, mfi)
}