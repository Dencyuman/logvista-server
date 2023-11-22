package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Dencyuman/logvista-server/config"
	api "github.com/Dencyuman/logvista-server/src/api"
	database "github.com/Dencyuman/logvista-server/src/database"
	util "github.com/Dencyuman/logvista-server/src/utils"
)

// @title LogVista API
// @version 0.1.11
// @description This is LogVista server.
// @BasePath /api/v1
func main() {
	seed := flag.Bool("seed", false, "Set to true to seed the database")
	migrate := flag.Bool("migrate", false, "Set to true to migrate the database")
	reset := flag.Bool("reset", false, "Set to true to reset the database")
	tmpl := flag.Bool("tmpl", false, "Set to true to run the server in Generate Template mode")
	flag.Parse()

	if *tmpl {
		staticDirPath := "./static/"
		path, err := util.FindFirstJSFile(staticDirPath)
		if err != nil {
			log.Fatal("Failed to find JS files:", err)
		}
		jsFile := filepath.Base(path)
		tmplPath := "./static/" + jsFile
		outputPath := "./static/assets/" + jsFile
		err = util.GenerateJSFileFromTemplate(tmplPath, outputPath)
		if err != nil {
			log.Fatal("Failed to generate JS file:", err)
		}
		return
	}

	db, err := initializeDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if *migrate {
		err = database.Migrate(db)
		if err != nil {
			log.Fatal("Failed to migrate the database:", err)
		}
		return
	}

	if *seed {
		err = database.Seed(db)
		if err != nil {
			log.Fatal("Failed to seed the database:", err)
		}
		return
	}

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

	router := api.SetupRouter(db)
	router.Run("0.0.0.0:" + config.AppConfig.ServerPort)

}
