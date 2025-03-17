package cmd

import (
	"github.com/tebruno99/nvms/nvdb"
	"gorm.io/gorm"
)

func InitGorm(db *gorm.DB) {
	db.AutoMigrate(&nvdb.File{})
}
