package service

import (
	domain "Application/internal/Domain"
	"context"
)

type ListingService interface {
	CreateListing(ctx context.Context, l *domain.Listing) (int64, error)
	GetListing(ctx context.Context, id int64) (*domain.Listing, error)
	MakeOffer(ctx context.Context, o *domain.Offer) (int64, error)
	AcceptOffer(ctx context.Context, offerID int64, agentID int64) error
	ScheduleAppointment(ctx context.Context, a *domain.Appointment) (int64, error)
}
