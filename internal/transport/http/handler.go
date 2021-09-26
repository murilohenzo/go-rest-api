package http

import (
	Entity "api/internal/comment/models"
	Comment "api/internal/comment/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
	Service *Comment.Service
}

// NewHandler - returns a pointer to Handler
func NewHandler(service *Comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("api/v1/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("api/v1/comments", h.PostComment).Methods("POST")
	h.Router.HandleFunc("api/v1/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("api/v1/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("api/v1/comments/{id}", h.DeleteComment).Methods("DELETE")
}

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_, err := fmt.Fprintf(writer, "Unable to parse UINT from ID")
		if err != nil {
			return 
		}
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		_, err := fmt.Fprintf(writer, "Error Retrieving Comment By ID")
		if err != nil {
			return 
		}
	}

	_, err = fmt.Fprintf(writer, "%+v", comment)
	if err != nil {
		return
	}
}

// GetAllComments - retrieves all comments from the component service
func (h *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		_, err := fmt.Fprintf(writer, "Failed to retrieve all comments")
		if err != nil {
			return 
		}
	}
	_, err = fmt.Fprintf(writer, "%+v", comments)
	if err != nil {
		return 
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	comment, err := h.Service.PostComment(Entity.Comment{
		Slug: "/",
	})
	if err != nil {
		_, err := fmt.Fprintf(writer, "Failed to post new comment")
		if err != nil {
			return
		}
	}
	_, err = fmt.Fprintf(writer, "%+v", comment)
	if err != nil {
		return
	}
}

// UpdateComment - updates a comment by ID with new comment info
func (h *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	comment, err := h.Service.UpdateComment(1, Entity.Comment{
		Slug: "/new",
	})
	if err != nil {
		_, err := fmt.Fprintf(writer, "Failed to update comment")
		if err != nil {
			return
		}
	}
	_, err = fmt.Fprintf(writer, "%+v", comment)
	if err != nil {
		return
	}
}

// DeleteComment - deletes a comment from the database by ID
func (h *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		_, err := fmt.Fprintf(writer, "Unable to parse UINT from ID")
		if err != nil {
			return
		}
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		_, err := fmt.Fprintf(writer, "Failed to delete comment")
		if err != nil {
			return 
		}
	}
}