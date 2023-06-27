package initializer

import (
	"github.com/aswinjithkukku/jwt-auth/models"
)

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
