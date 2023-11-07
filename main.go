package main

import (
	"flag"
	"log"

	api "github.com/Dencyuman/logvista-server/src/api"
	database "github.com/Dencyuman/logvista-server/src/database"
)

// @title LogVista API
// @version 0.1.4
// @description This is LogVista server.
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	seed := flag.Bool("seed", false, "Set to true to seed the database")
	flag.Parse()

	db, err := initializeDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	if *seed {
		err = database.Seed(db)
		if err != nil {
			log.Fatal("Failed to seed the database:", err)
		}
	} else {
		router := api.SetupRouter(db)
		router.Run(":8080")
	}

}
