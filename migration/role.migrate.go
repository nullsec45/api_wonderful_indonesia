package migration

import (
	"api/database"
	"api/models"
)

func RoleMigrate() {
	db := database.Init()
	var role models.Role
	roleTable := db.Migrator().HasTable(role)

	if !roleTable {
		db.AutoMigrate(role)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
