package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func RoleMigrate() {
	db := database.ConnectDatabase()
	var role models.Role
	roleTable := db.Migrator().HasTable(role)

	if !roleTable {
		db.AutoMigrate(role)
		fmt.Println("Berhasil membuat table daftar_role")
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
