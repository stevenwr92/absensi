package migrations

import (
	"github.com/stevenwr92/absensi/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	// Drop existing tables if they exist
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable(&models.Attendance{})

	// AutoMigrate to create tables
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Attendance{})
}
