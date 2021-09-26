package database

import (
	Entity "api/internal/comment/models"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	var comment Entity.Comment
	if result := db.AutoMigrate(&comment); result.Error != nil {
		return result.Error
	}
	return nil
}