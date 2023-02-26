package migration

import (
	"api/database"
	"api/models"
)

func PakaianAdatMigrate() {
	db := database.Init()
	var pakaianAdat models.PakaianAdat

	pakaianAdatTable := db.Migrator().HasTable(pakaianAdat)

	if !pakaianAdatTable {
		db.AutoMigrate(pakaianAdat)
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
