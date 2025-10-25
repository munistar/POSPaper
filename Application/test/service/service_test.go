package service_test

import (
	domain "Application/internal/Domain"
	repo "Application/internal/Repository"
	service "Application/internal/service"
	"context"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	user := &domain.User{
		Username:  "beforeupdate",
		Firstname: "Before",
		Email:     "before@example.com",
	}

	created, _ := s.Create(ctx, user)

	updated := &domain.User{
		Username:  "afterupdate",
		Firstname: "After",
		Email:     "after@example.com",
	}

	result, err := s.Update(ctx, created.ID, updated)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if result.Username != "afterupdate" {
		t.Errorf("Expected Username 'afterupdate', got %s", result.Username)
	}
	if result.Firstname != "After" {
		t.Errorf("Expected Firstname 'After', got %s", result.Firstname)
	}
}

func TestUserService_Update_NotFound(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	user := &domain.User{Username: "test"}
	_, err := s.Update(ctx, "non-existent", user)

	if err != repo.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}

func TestUserService_Delete(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	user := &domain.User{Username: "deletethis"}
	created, _ := s.Create(ctx, user)

	err := s.Delete(ctx, created.ID)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	_, err = s.GetByID(ctx, created.ID)
	if err != repo.ErrNotFound {
		t.Error("User should be deleted")
	}
}

/*
	func TestUserService_Delete_NotFound(t *testing.T) {
		r := repo.NewMemoryUserRepo()
		s := service.NewUserService(r)
		ctx := context.Background()

		err := s.Delete(ctx, "non-existent")
		if err != repo.ErrNotFound {
			t.Errorf("Expected ErrNotFound, got %v", err)
		}
	}Username:  "serviceuser",

Firstname: "Service",
Lastname:  "User",
Email:     "service@example.com",
Password:  "hashed_password",
}

created, err := s.Create(ctx, user)
if err != nil {
t.Fatalf("Create failed: %v", err)
}

if created.ID == "" {
t.Error("Expected ID to be set")
}
if created.Username != "serviceuser" {
t.Errorf("Expected Username 'serviceuser', got %s", created.Username)
}
}
*/
func TestUserService_GetByID(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	user := &domain.User{
		Username: "findviaservice",
		Email:    "find@example.com",
	}

	created, _ := s.Create(ctx, user)

	found, err := s.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}

	if found.Username != "findviaservice" {
		t.Errorf("Expected Username 'findviaservice', got %s", found.Username)
	}
}

func TestUserService_GetByID_NotFound(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	_, err := s.GetByID(ctx, "non-existent")
	if err != repo.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}

func TestUserService_List(t *testing.T) {
	r := repo.NewMemoryUserRepo()
	s := service.NewUserService(r)
	ctx := context.Background()

	users := []*domain.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
		{Username: "user3", Email: "user3@example.com"},
	}

	for _, u := range users {
		s.Create(ctx, u)
	}

	list, err := s.List(ctx)
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if len(list) != 3 {
		t.Errorf("Expected 3 users, got %d", len(list))
	}
}
