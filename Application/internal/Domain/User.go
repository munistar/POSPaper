package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"Email" db:"Email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
