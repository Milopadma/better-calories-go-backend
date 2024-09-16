package days

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

func (h *Handler) CreateDay(c echo.Context) error {
	d := new(Day)
	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := h.DB.QueryRow("INSERT INTO days (user_id, date) VALUES ($1, $2) RETURNING id", d.UserID, d.Date).Scan(&d.ID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating day"})
	}

	return c.JSON(http.StatusCreated, d)
}

func (h *Handler) GetUserDays(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	rows, err := h.DB.Query("SELECT id, date FROM days WHERE user_id = $1", userID)
	if err != nil {
		fmt.Printf("DEBUG Database error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching user days"})
	}
	defer rows.Close()

	var days []Day
	for rows.Next() {
		var day Day
		if err := rows.Scan(&day.ID, &day.Date); err != nil {
			fmt.Printf("DEBUG Row scan error: %v\n", err)
			continue
		}
		day.UserID = userID
		days = append(days, day)
	}

	return c.JSON(http.StatusOK, days)
}