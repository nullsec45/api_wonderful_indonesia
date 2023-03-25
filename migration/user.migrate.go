package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func UserMigrate() {
	db := database.ConnectDatabase()

	var user models.User
	userTable := db.Migrator().HasTable(user)
	if !userTable {
		db.AutoMigrate(user)
		fmt.Println("Berhasil membuat table daftar_user")
	}
}
