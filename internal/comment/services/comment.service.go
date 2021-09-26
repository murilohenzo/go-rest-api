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

// GetComment - retrieves comment by their ID from the database
func (s *Service) GetComment(ID uint) (Entity.Comment, error) {
	var comment Entity.Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Entity.Comment{}, result.Error
	}

	return comment, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentsBySlug(slug string) ([]Entity.Comment, error) {
	var comments []Entity.Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Entity.Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Entity.Comment) (Entity.Comment, error) {
	if result := s.DB.Create(&comment); result.Error != nil {
		return Entity.Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ID uint, newComment Entity.Comment) (Entity.Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Entity.Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Entity.Comment{}, result.Error
	}
	return comment, nil

}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Entity.Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - retrieves all comments from the database
func (s *Service) GetAllComments() ([]Entity.Comment, error) {
	var comments []Entity.Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}