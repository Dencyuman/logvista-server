package main

import (
	"log"

	api "github.com/Dencyuman/logvista-server/src/api"
	database "github.com/Dencyuman/logvista-server/src/database"
)

// @title LogVista API
// @version 0.1.3
// @description This is LogVista server.
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	db, err := initializeDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}
	router := api.SetupRouter(db)

	// fmt.Println(db) // If you don't need this, comment or remove
	router.Run(":8080")
}
