package migrations

import (
	"github.com/stevenwr92/absensi/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable(&models.Attendance{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Attendance{})
}
