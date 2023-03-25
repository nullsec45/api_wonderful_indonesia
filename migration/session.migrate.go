package migration

import (
	"api/database"
	"api/models"
	"fmt"
)

func SessionMigrate() {
	db := database.ConnectDatabase()
	var session models.Session
	sessionTable := db.Migrator().HasTable(session)

	if !sessionTable {
		db.AutoMigrate(session)
		fmt.Println("Berhasil membuat table session_token")
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}
