package main

/*
import (
	"log"
	"net/http"

	repo "Application/internal/Repostiory"
	service "Application/internal/service"
	transport "Application/internal/transport/http"
)

func main() {
	// repo -> service -> handler
	propertyRepo := repo.NewMemoryPropertyRepo()
	propertyService := service.NewPropertyService(propertyRepo)
	propertyHandler := transport.NewPropertyHandler(propertyService)

	router := transport.NewRouter(propertyHandler)

	log.Println("🚀 Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
*/

import (
	repo "Application/internal/Repostiory"
	transport "Application/internal/Transport/http"
	service "Application/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	propertyRepo := repo.NewMemoryPropertyRepo()
	propertyService := service.NewPropertyService(propertyRepo)
	propertyHandler := transport.NewPropertyHandler(propertyService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API routes
	r.Route("/api/v1/properties", func(rt chi.Router) {
		rt.Mount("/", propertyHandler.Routes())
	})

	// Serve frontend
	fs := http.FileServer(http.Dir("./web"))
	r.Handle("/*", fs)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Println("API running at http://localhost:8080/api/v1/properties")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":8080", r))

}
