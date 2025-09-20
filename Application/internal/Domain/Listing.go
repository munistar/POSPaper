package domain

import "time"

type ListingStatus string

const (
	ListingActive   ListingStatus = "active"
	ListingPending  ListingStatus = "pending"
	ListingSold     ListingStatus = "sold"
	ListingArchived ListingStatus = "archived"
)

type Listing struct {
	ID          int64         `json:"id" db:"id"`
	PropertyID  int64         `json:"property_id" db:"property_id"`
	AgentID     int64         `json:"agent_id" db:"agent_id"`
	AskingPrice int64         `json:"asking_price" db:"asking_price"` // cents
	Status      ListingStatus `json:"status" db:"status"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}
