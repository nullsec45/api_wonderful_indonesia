package migration

import (
	"api/database"
	"api/models"
)

func KabupatenMigrate() {
	db := database.Init()
	var kabupaten models.Kabupaten
	kabupatenTable := db.Migrator().HasTable(kabupaten)

	if !kabupatenTable {
		db.AutoMigrate(kabupaten)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
