package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func KabupatenMigrate() {
	db := database.ConnectDatabase()
	var kabupaten models.Kabupaten
	kabupatenTable := db.Migrator().HasTable(kabupaten)

	if !kabupatenTable {
		db.AutoMigrate(kabupaten)
		fmt.Println("Berhasil membuat table daftar_kabupaten")
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
