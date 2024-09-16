package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/milopadma/better-calories-go-backend/internal/database"
	"github.com/milopadma/better-calories-go-backend/internal/days"
	"github.com/milopadma/better-calories-go-backend/internal/fooditems"
	"github.com/milopadma/better-calories-go-backend/internal/mealfooditems"
	"github.com/yourusername/yourproject/internal/meals"
	"github.com/yourusername/yourproject/internal/server"
	"github.com/yourusername/yourproject/internal/users"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&users.User{}, &days.Day{}, &meals.Meal{}, &fooditems.FoodItem{}, &mealfooditems.MealFoodItem{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	s := server.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := s.Start(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
