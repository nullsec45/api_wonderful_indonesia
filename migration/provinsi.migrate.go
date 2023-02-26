package migration

import (
	"api/database"
	"api/models"
)

func ProvinsiMigrate() {
	db := database.Init()
	var provinsi models.Provinsi
	provinsiTable := db.Migrator().HasTable(provinsi)

	if !provinsiTable {
		db.AutoMigrate(provinsi)
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
