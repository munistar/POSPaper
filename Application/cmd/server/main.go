package main

import (
	repo "Application/internal/Repository"
	service "Application/internal/service"
	transport "Application/internal/transport/http"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize repositories
	propertyRepo := repo.NewMemoryPropertyRepo()
	userRepo := repo.NewMemoryUserRepo()

	// Initialize services
	propertyService := service.NewPropertyService(propertyRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	propertyHandler := transport.NewPropertyHandler(propertyService)
	userHandler := transport.NewUserHandler(userService)

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if req.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, req)
		})
	})

	// API routes
	r.Route("/api/v1", func(rt chi.Router) {
		rt.Mount("/properties", propertyHandler.Routes())
		rt.Mount("/users", userHandler.Routes())
	})

	// Serve frontend static files
	fs := http.FileServer(http.Dir("./web"))
	r.Handle("/*", fs)

	log.Println("🚀 Server running at http://localhost:8080/login.html")
	log.Println("📍 API Endpoints:")
	log.Println("   Properties: http://localhost:8080/api/v1/properties")
	log.Println("   Users: http://localhost:8080/api/v1/users")
	log.Fatal(http.ListenAndServe(":8080", r))
}
