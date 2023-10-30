package main

import (
	"log"

	"github.com/Oluwaseun241/swiftclip-backend/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  // database connection
  db.Init()

  err := app.Listen(":3000")
  if err != nil {
    log.Fatalf("Error starting server: %v", err)
  }
}
