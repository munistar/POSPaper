package service

import (
	domain "Application/internal/Domain"
	repo "Application/internal/Repository"
	"context"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(ctx context.Context, p *domain.User) (*domain.User, error) {
	return s.repo.Create(ctx, p)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]*domain.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, id string, p *domain.User) (*domain.User, error) {
	return s.repo.Update(ctx, id, p)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
