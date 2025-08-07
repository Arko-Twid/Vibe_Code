package main

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
	Price       float64
	ImageURL    string
	Category    string
	Tags        []string `gorm:"type:text[]"`
	Stock       int
	CreatedAt   time.Time
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email     string    `gorm:"unique"`
	Password  string
	Role      string // customer, admin
	CreatedAt time.Time
}
