package migration

import (
	"api/database"
	"api/models"
)

func UserMigrate() {
	db := database.Init()

	var user models.User
	userTable := db.Migrator().HasTable(user)
	if !userTable {
		db.AutoMigrate(user)
	}
}
