package main

import (
	"fmt"
	"os"

	"github.com/supardi98/golang-gorm-postgres/cmd"
	"github.com/supardi98/golang-gorm-postgres/routes"
)

func main() {
	// Argumens function
	args := os.Args
	if len(args) > 1 {
		if args[1] == "migrate" {
			fmt.Println("Migrating")

			cmd.Migrate()

			os.Exit(0)
		} else if args[1] == "seed" {
			fmt.Println("Seeding")

			cmd.Seed()

			os.Exit(0)
		}
	}
	routes.Init()
}
