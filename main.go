package main

import (
	"api_wonderful_indonesia/database"
	"api_wonderful_indonesia/routes"
)

func main() {
	database.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))

}
