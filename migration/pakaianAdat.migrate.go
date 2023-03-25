package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func PakaianAdatMigrate() {
	db := database.ConnectDatabase()
	var pakaianAdat models.PakaianAdat

	pakaianAdatTable := db.Migrator().HasTable(pakaianAdat)

	if !pakaianAdatTable {
		db.AutoMigrate(pakaianAdat)
		fmt.Println("Berhasil membuat table daftar_pakaian_adat")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
