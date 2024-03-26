package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/Dencyuman/logvista-server/config"
	"github.com/Dencyuman/logvista-server/src/api"
	"github.com/Dencyuman/logvista-server/src/background"
	"github.com/Dencyuman/logvista-server/src/database"
	"github.com/Dencyuman/logvista-server/src/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// @title Logvista API
// @version 0.1.13
// @description This is Logvista server.
// @BasePath /api/v1
func main() {
	seed := flag.Bool("seed", false, "Set to true to seed the database")
	migrate := flag.Bool("migrate", false, "Set to true to migrate the database")
	reset := flag.Bool("reset", false, "Set to true to reset the database")
	flag.Parse()

	// Generate JS file from template
	staticDirPath := "./static/"
	path, err := utils.FindFirstJSFile(staticDirPath)
	if err != nil {
		log.Fatal("Failed to find JS files:", err)
	}
	jsFile := filepath.Base(path)
	tmplPath := "./static/" + jsFile
	outputPath := "./static/assets/" + jsFile
	err = utils.GenerateJSFileFromTemplate(tmplPath, outputPath)
	if err != nil {
		log.Fatal("Failed to generate JS file:", err)
	}

	// Initialize the database
	db, err := initializeDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the database (if necessary)
	if *migrate {
		err = database.Migrate(db)
		if err != nil {
			log.Fatal("Failed to migrate the database:", err)
		}
		return
	}

	// Seed the database (if necessary)
	if *seed {
		err = database.Seed(db)
		if err != nil {
			log.Fatal("Failed to seed the database:", err)
		}
		return
	}

	// Reset the database (if necessary)
	if *reset {
		reader := bufio.NewReader(os.Stdin)
		log.Println("Are you sure you want to reset the database? (y/n):")
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read response:", err)
		}

		response = strings.TrimSpace(response)
		if strings.ToLower(response) == "y" {
			err = database.ResetTables(db)
			if err != nil {
				log.Fatal("Failed to reset the database:", err)
			}
			log.Println("Database reset successfully.")
		} else {
			log.Println("Database reset cancelled.")
		}
		return
	}

	background.SetupServerChecker(db)
	// Start the server
	fmt.Printf("Starting server on port %s...\n", config.AppConfig.ServerPort)
	router := api.SetupRouter(db)
	router.Run("0.0.0.0:" + config.AppConfig.ServerPort)
}
