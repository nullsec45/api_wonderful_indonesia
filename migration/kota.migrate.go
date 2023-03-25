package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func KotaMigrate() {
	db := database.ConnectDatabase()
	var kota models.Kota

	kotaTable := db.Migrator().HasTable(kota)

	if !kotaTable {
		db.AutoMigrate(kota)
		fmt.Println("Berhasil membuat table daftar_kota")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
