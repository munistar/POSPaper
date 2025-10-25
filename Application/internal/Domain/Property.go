package domain

import "time"

type Property struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"`
	SizeSqm     int       `json:"size_sqm" db:"size_sqm"`
	Bedrooms    int       `json:"bedrooms" db:"bedrooms"`
	Bathrooms   int       `json:"bathrooms" db:"bathrooms"`
	Address     Address   `json:"address" db:"-"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Owner       User      `json:"owner" db:"owner"`
}

type Address struct {
	Street     string `json:"street" db:"street"`
	City       string `json:"city" db:"city"`
	PostalCode string `json:"postal_code" db:"postal_code"`
	Country    string `json:"country" db:"country"`
}
