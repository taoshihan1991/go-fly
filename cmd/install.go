package cmd

import (
	"github.com/spf13/cobra"
	"goflylivechat/models"
	"goflylivechat/tools"
	"log"
	"os"
	"strings"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Initialize database and import data", // More precise description
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func install() {
	// Check if already installed
	if ok, _ := tools.IsFileNotExist("./install.lock"); !ok {
		log.Println("Please remove ./install.lock file to reinstall")
		os.Exit(1)
	}

	// Verify required files exist
	sqlFile := "import.sql"
	dataExists, _ := tools.IsFileExist(sqlFile)
	if !dataExists {
		log.Println("Configuration file config/mysql.json or database import file import.sql not found")
		os.Exit(1)
	}

	// Execute SQL statements
	sqls, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Printf("Failed to read SQL file %s: %v\n", sqlFile, err)
		os.Exit(1)
	}
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := models.Execute(sql)
		if err != nil {
			log.Printf("SQL execution failed: %s\nError: %v\n", sql, err)
			log.Println("Database initialization failed - please check SQL statements")
			os.Exit(1)
		}
		log.Printf("Executed successfully: %s\n", sql)
	}

	// Create installation lock file
	installFile, err := os.OpenFile("./install.lock", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("Failed to create lock file: %v\n", err)
		os.Exit(1)
	}
	defer installFile.Close()

	_, err = installFile.WriteString("gofly live chat installation complete")
	if err != nil {
		log.Printf("Failed to write lock file: %v\n", err)
		os.Exit(1)
	}

	log.Println("Database initialization completed successfully")
}
