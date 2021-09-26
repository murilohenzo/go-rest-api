package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/v1/health", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "I am alive!")
		if err != nil {
			return
		}
	})
}