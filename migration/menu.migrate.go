package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func MenuMigrate() {
	db := database.ConnectDatabase()
	var menu models.Menu
	menuTable := db.Migrator().HasTable(menu)

	if !menuTable {
		db.AutoMigrate(menu)
		fmt.Println("Berhasil membuat table daftar_menu")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
