package database

import (
	"gorm.io/gorm"
	"test-backend-developer-sagala/domain"
)

func MigrateTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&domain.User{}, &domain.Article{})

	return err
}
