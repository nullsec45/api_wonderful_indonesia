package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func UserRoleMigrate() {
	db := database.ConnectDatabase()
	var user_role models.UserRole
	userRoleTable := db.Migrator().HasTable(user_role)

	if !userRoleTable {
		db.AutoMigrate(user_role)
		fmt.Println("Berhasil membuat table daftar_user_role")
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
