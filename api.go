package main

import (
	"api_wonderful_indonesia/database"
	"api_wonderful_indonesia/routes"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	database.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))

	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
