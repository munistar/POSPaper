package domain

import "time"

type OfferStatus string

const (
	OfferPending  OfferStatus = "pending"
	OfferAccepted OfferStatus = "accepted"
	OfferRejected OfferStatus = "rejected"
)

type Offer struct {
	ID         int64       `json:"id" db:"id"`
	ListingID  int64       `json:"listing_id" db:"listing_id"`
	CustomerID int64       `json:"customer_id" db:"customer_id"`
	Amount     int64       `json:"amount" db:"amount"` // cents
	Status     OfferStatus `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
}

type AppointmentStatus string

type Appointment struct {
	ID         int64             `json:"id" db:"id"`
	ListingID  int64             `json:"listing_id" db:"listing_id"`
	AgentID    int64             `json:"agent_id" db:"agent_id"`
	CustomerID int64             `json:"customer_id" db:"customer_id"`
	StartTime  time.Time         `json:"start_time" db:"start_time"`
	EndTime    time.Time         `json:"end_time" db:"end_time"`
	Notes      string            `json:"notes" db:"notes"`
	Status     AppointmentStatus `json:"status" db:"status"`
}

type Contract struct {
	ID       int64     `json:"id" db:"id"`
	OfferID  int64     `json:"offer_id" db:"offer_id"`
	SignedAt time.Time `json:"signed_at" db:"signed_at"`
	FileURL  string    `json:"file_url" db:"file_url"`
}
