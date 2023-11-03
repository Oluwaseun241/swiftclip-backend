package db

import (
	"fmt"
	"os"

	"github.com/Oluwaseun241/swiftclip-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
  godotenv.Load()
  var (
    dbname = os.Getenv("DB_NAME")
    dbuser = os.Getenv("DB_USER")
    dbpassword = os.Getenv("DB_PASSWORD")
    dbhost = os.Getenv("DB_HOST")
    uri = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432", dbuser, dbname, dbpassword, dbhost)
  )
  
  db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
  if err != nil {
    panic("Failed to connect to DB")
  }
  fmt.Printf("DB connected")

  DB = db
  db.AutoMigrate(&models.User{})
  fmt.Printf("DB migrated")
}
