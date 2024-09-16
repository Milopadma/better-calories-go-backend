package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	result := h.DB.Create(u)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating user"})
	}

	u.PasswordHash = "" // Don't send password hash back
	return c.JSON(http.StatusCreated, u)
}