package Repository

import (
	domain "Application/internal/Domain"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type PropertyRepository interface {
	Create(ctx context.Context, p *domain.Property) (*domain.Property, error)
	GetByID(ctx context.Context, id string) (*domain.Property, error)
	List(ctx context.Context) ([]*domain.Property, error)
	Update(ctx context.Context, id string, p *domain.Property) (*domain.Property, error)
	Delete(ctx context.Context, id string) error
}

type MemoryPropertyRepo struct {
	mu         sync.RWMutex
	properties map[string]*domain.Property
}

func NewMemoryPropertyRepo() *MemoryPropertyRepo {
	return &MemoryPropertyRepo{
		properties: make(map[string]*domain.Property),
	}
}

func (r *MemoryPropertyRepo) Create(ctx context.Context, p *domain.Property) (*domain.Property, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = p.CreatedAt
	r.properties[p.ID] = p
	return p, nil
}

func (r *MemoryPropertyRepo) GetByID(ctx context.Context, id string) (*domain.Property, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.properties[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (r *MemoryPropertyRepo) List(ctx context.Context) ([]*domain.Property, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	props := make([]*domain.Property, 0, len(r.properties))
	for _, p := range r.properties {
		props = append(props, p)
	}
	return props, nil
}

func (r *MemoryPropertyRepo) Update(ctx context.Context, id string, p *domain.Property) (*domain.Property, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.properties[id]
	if !ok {
		return nil, ErrNotFound
	}
	p.ID = id
	p.CreatedAt = existing.CreatedAt
	p.UpdatedAt = time.Now().UTC()
	r.properties[id] = p
	return p, nil
}

func (r *MemoryPropertyRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.properties[id]; !ok {
		return ErrNotFound
	}
	delete(r.properties, id)
	return nil
}
