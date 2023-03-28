package initializers

import "auth-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
