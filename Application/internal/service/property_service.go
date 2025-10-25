package service

import (
	domain "Application/internal/Domain"
	repo "Application/internal/Repository"
	"context"
)

type PropertyService struct {
	repo repo.PropertyRepository
}

func NewPropertyService(r repo.PropertyRepository) *PropertyService {
	return &PropertyService{repo: r}
}

func (s *PropertyService) Create(ctx context.Context, p *domain.Property) (*domain.Property, error) {
	return s.repo.Create(ctx, p)
}

func (s *PropertyService) GetByID(ctx context.Context, id string) (*domain.Property, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PropertyService) List(ctx context.Context) ([]*domain.Property, error) {
	return s.repo.List(ctx)
}

func (s *PropertyService) Update(ctx context.Context, id string, p *domain.Property) (*domain.Property, error) {
	return s.repo.Update(ctx, id, p)
}

func (s *PropertyService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
