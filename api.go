package main

import (
	"api/database"
	"api/migration"
	"api/routes"
	"fmt"

	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "api",
		Usage: "Perintah ini digunakan untuk manajemen api",
		Action: func(cCtx *cli.Context) error {
			args := cCtx.Args().Get(0)
			//fmt.Printf("Hello %q", cCtx.Args().Get(0))

			switch args {
			case "migrate":
				migration.Migrate()
			case "run":
				database.ConnectDatabase()
				routes.Init()
			default:
				fmt.Println("Selamat Datang di API Wonderful Of Indonesia coy")
				fmt.Println("API ini berisikan tentang indahnya keberagaman yang ada di Indonesia.")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
