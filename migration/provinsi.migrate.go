package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func ProvinsiMigrate() {
	db := database.ConnectDatabase()
	var provinsi models.Provinsi
	provinsiTable := db.Migrator().HasTable(provinsi)

	if !provinsiTable {
		db.AutoMigrate(provinsi)
		fmt.Println("Berhasil membuat table daftar_provinsi")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
