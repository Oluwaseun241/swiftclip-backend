package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model `gorm:"tables:users,alias:u"`

  ID  int64 `gorm:"primaryKey,autoIncrement"`
  Email string  `json:"string"`
  Password  string
}

