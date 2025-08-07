package main

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Product represents a stationary item in the marketplace
// Uses Postgres uuid and text[] types
// GORM auto-migrates these fields for Postgres

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Category    string    `json:"category"`
	Tags        []string  `gorm:"type:text[]" json:"tags"` // Postgres array
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

// User represents a customer or admin
// Uses Postgres uuid type for ID

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`    // omit from JSON responses
	Role      string    `json:"role"` // customer, admin
	CreatedAt time.Time `json:"created_at"`
}

// Password helpers for User
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
