package services

import (
	Entity "api/internal/comment/models"
	"github.com/jinzhu/gorm"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}



// ServiceComment - the interface for our comment service
type ServiceComment interface {
	GetComment(ID uint) (Entity.Comment, error)
	GetCommentsBySlug(slug string) ([]Entity.Comment, error)
	PostComment(comment Entity.Comment) (Entity.Comment, error)
	UpdateComment(ID uint, newComment Entity.Comment) (Entity.Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Entity.Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}