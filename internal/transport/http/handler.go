package http

import (
	Entity "api/internal/comment/models"
	Comment "api/internal/comment/services"
	"encoding/json"
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

// Response - an object to store response from our api
type Response struct {
	Message string
	Error string
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

	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods("DELETE")
}

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Unable to parse UINT from ID", err, http.StatusBadRequest)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(writer, "Error Retrieving Comment By ID", err, http.StatusNotFound)
		return
	}

	if err := sendOkResponse(writer, comment, http.StatusOK); err != nil {
		panic(err)
	}
}

// GetAllComments - retrieves all comments from the component service
func (h *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(writer, "Failed to retrieve all comments", err, http.StatusNotFound)
		return
	}
	if err := sendOkResponse(writer, comments, http.StatusOK); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	var comment Entity.Comment
	if err := json.NewDecoder(request.Body).Decode(&comment); err != nil {
		sendErrorResponse(writer, "Failed to decode JSON Body", err, http.StatusBadRequest)
		return
	}

	createdComment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(writer, "Failed to post new comment", err, http.StatusBadRequest)
		return
	}
	if err := sendOkResponse(writer, createdComment, http.StatusCreated); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment by ID with new comment info
func (h *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Unable to parse UINT from ID", err, http.StatusBadRequest)
		return
	}

	var comment Entity.Comment
	if err := json.NewDecoder(request.Body).Decode(&comment); err != nil {
		sendErrorResponse(writer, "Failed to decode JSON Body", err, http.StatusBadRequest)
		return
	}

	updatedComment, err := h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		sendErrorResponse(writer, "Failed to update comment", err, http.StatusNotFound)
		return
	}

	if err := sendOkResponse(writer, updatedComment, http.StatusCreated); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment from the database by ID
func (h *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Unable to parse UINT from ID", err, http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(writer, "Failed to delete comment", err, http.StatusBadRequest)
		return
	}

	if err := sendOkResponse(writer, Response{
		Message: "Successfully delete comment",
	}, http.StatusOK); err != nil {
		panic(err)
	}
}

func sendOkResponse(writer http.ResponseWriter, resp interface{}, statusCode int) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(resp)
}

func sendErrorResponse(writer http.ResponseWriter, message string, err error, statusCode int) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(statusCode)
	if err := json.NewEncoder(writer).Encode(Response{
		Message: message,
		Error: err.Error(),
	}); err != nil {
		panic(err)
	}
}