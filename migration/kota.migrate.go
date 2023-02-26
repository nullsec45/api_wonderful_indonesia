package migration

import (
	"api/database"
	"api/models"
)

func KotaMigrate() {
	db := database.Init()
	var kota models.Kota

	kotaTable := db.Migrator().HasTable(kota)

	if !kotaTable {
		db.AutoMigrate(kota)
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
