package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func RoleMenuMigrate() {
	db := database.ConnectDatabase()
	var role_menu models.RoleMenu
	roleMenuTable := db.Migrator().HasTable(role_menu)

	if !roleMenuTable {
		db.AutoMigrate(role_menu)
		fmt.Println("Berhasil membuat table daftar_role_menu")
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
