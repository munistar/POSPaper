package domain

import (
	domain "Application/internal/Domain"
	"testing"
	"time"
)

func TestPropertyStructure(t *testing.T) {
	now := time.Now()

	addr := domain.Address{
		Street:     "123 Main St",
		City:       "Vienna",
		PostalCode: "1010",
		Country:    "Austria",
	}

	user := domain.User{
		ID:        "user-1",
		Username:  "testuser",
		Firstname: "Test",
		Lastname:  "User",
		Email:     "test@example.com",
		CreatedAt: now,
	}

	prop := domain.Property{
		ID:          "prop-1",
		Title:       "Beautiful Apartment",
		Description: "A lovely place",
		Type:        "apartment",
		SizeSqm:     100,
		Bedrooms:    2,
		Bathrooms:   1,
		Address:     addr,
		CreatedAt:   now,
		UpdatedAt:   now,
		Owner:       user,
	}

	if prop.ID != "prop-1" {
		t.Errorf("Expected ID 'prop-1', got %s", prop.ID)
	}
	if prop.Title != "Beautiful Apartment" {
		t.Errorf("Expected Title 'Beautiful Apartment', got %s", prop.Title)
	}
	if prop.SizeSqm != 100 {
		t.Errorf("Expected SizeSqm 100, got %d", prop.SizeSqm)
	}
	if prop.Address.City != "Vienna" {
		t.Errorf("Expected City 'Vienna', got %s", prop.Address.City)
	}
	if prop.Owner.Username != "testuser" {
		t.Errorf("Expected Owner Username 'testuser', got %s", prop.Owner.Username)
	}
}

func TestAddressStructure(t *testing.T) {
	addr := domain.Address{
		Street:     "456 Oak Ave",
		City:       "Salzburg",
		PostalCode: "5020",
		Country:    "Austria",
	}

	if addr.Street != "456 Oak Ave" {
		t.Errorf("Expected Street '456 Oak Ave', got %s", addr.Street)
	}
	if addr.PostalCode != "5020" {
		t.Errorf("Expected PostalCode '5020', got %s", addr.PostalCode)
	}
}
