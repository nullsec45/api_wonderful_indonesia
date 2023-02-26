package main

import (
	"api_wonderful_indonesia/database"
	"api_wonderful_indonesia/models"
)

func Migrate() bool {
	db := database.Init()

	var user models.User
	var kabupaten models.Kabupaten
	var kota models.Kota
	var pakaianAdat models.PakaianAdat
	var provinsi models.Provinsi
	var role models.Role

	//db.AutoMigrate(&models.User{}, &models.Kabupaten{}, &models.Kota{}, &models.PakaianAdat{}, &models.Provinsi{})
	userTable := db.Migrator().HasTable(user)
	kabupatenTable := db.Migrator().HasTable(kabupaten)
	kotaTable := db.Migrator().HasTable(kota)
	pakaianAdatTable := db.Migrator().HasTable(pakaianAdat)
	provinsiTable := db.Migrator().HasTable(provinsi)
	roleTable := db.Migrator().HasTable(role)

	if !userTable && !kabupatenTable && !kotaTable && !pakaianAdatTable && !provinsiTable && !roleTable {
		db.AutoMigrate(user, kabupaten, kota, pakaianAdat, provinsi, roleTable)
		return true
	} else {
		//fmt.Println("Table sudah ada!")
		//db.close()
		return false
	}

}
