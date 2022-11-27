package initializers

import "github.com/jaimeiherrera/jwt_auth_golang/models"

func Migrate_db() {
	DB.AutoMigrate(&models.User{})
}
