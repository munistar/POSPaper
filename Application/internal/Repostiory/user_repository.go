package Repostiory

import (
	domain "Application/internal/Domain"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

//var ErrNotFound = errors.New("not found")

type UserRepository interface {
	Create(ctx context.Context, p *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, id string, p *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type MemoryUserRepo struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{
		users: make(map[string]*domain.User),
	}
}

func (r *MemoryUserRepo) Create(ctx context.Context, p *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	r.users[p.ID] = p
	return p, nil
}

func (r *MemoryUserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (r *MemoryUserRepo) List(ctx context.Context) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	props := make([]*domain.User, 0, len(r.users))
	for _, p := range r.users {
		props = append(props, p)
	}
	return props, nil
}

func (r *MemoryUserRepo) Update(ctx context.Context, id string, p *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	p.ID = id
	p.CreatedAt = existing.CreatedAt
	r.users[id] = p
	return p, nil
}

func (r *MemoryUserRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[id]; !ok {
		return ErrNotFound
	}
	delete(r.users, id)
	return nil
}
