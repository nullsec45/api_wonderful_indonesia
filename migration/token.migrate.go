package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func TokenMigrate() {
	db := database.ConnectDatabase()
	var token models.Token
	tokenTable := db.Migrator().HasTable(token)

	if !tokenTable {
		db.AutoMigrate(token)
		fmt.Println("Berhasil membuat table token_api")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
