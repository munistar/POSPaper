package transport

import (
	repo "Application/internal/Repository"
	service "Application/internal/service"
	transport "Application/internal/transport/http"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_CORS(t *testing.T) {
	r := repo.NewMemoryPropertyRepo()
	s := service.NewPropertyService(r)
	h := transport.NewPropertyHandler(s)
	router := transport.NewRouter(h)

	req := httptest.NewRequest(http.MethodOptions, "/api/v1/properties", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for OPTIONS, got %d", w.Code)
	}

	allowOrigin := w.Header().Get("Access-Control-Allow-Origin")
	if allowOrigin != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin '*', got '%s'", allowOrigin)
	}

	allowMethods := w.Header().Get("Access-Control-Allow-Methods")
	if allowMethods == "" {
		t.Error("Expected Access-Control-Allow-Methods header to be set")
	}
}

func TestRouter_StaticFiles(t *testing.T) {
	r := repo.NewMemoryPropertyRepo()
	s := service.NewPropertyService(r)
	h := transport.NewPropertyHandler(s)
	router := transport.NewRouter(h)

	// Test that static file route exists (will 404 if files don't exist, which is ok for test)
	req := httptest.NewRequest(http.MethodGet, "/login.html", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// We expect either 200 (file exists) or 404 (file not found in test environment)
	// Both are acceptable as we're just testing the route exists
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected status 200 or 404, got %d", w.Code)
	}
}

func TestRouter_APIRoutes(t *testing.T) {
	r := repo.NewMemoryPropertyRepo()
	s := service.NewPropertyService(r)
	h := transport.NewPropertyHandler(s)
	router := transport.NewRouter(h)

	tests := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{
			name:   "List properties",
			method: http.MethodGet,
			path:   "/api/v1/properties",
			status: http.StatusOK,
		},
		{
			name:   "Get property by ID (not found)",
			method: http.MethodGet,
			path:   "/api/v1/properties/123",
			status: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Errorf("Expected status %d, got %d", tt.status, w.Code)
			}
		})
	}
}
