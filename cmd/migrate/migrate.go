package main

import (
	"database/sql"
	"fmt"
	migrations "github.com/MrRytis/apple-business-manager/internal/migration"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func main() {
	color.Green("Running migration up")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open: ", err)
	}

	defer db.Close()

	goose.SetBaseFS(migrations.EmbedMigrationsFs)

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal("Failed to set dialect: ", err)
	}

	current, err := goose.GetDBVersion(db)
	if err != nil {
		log.Fatal("Failed to get current version: ", err)
	}

	color.Green(fmt.Sprintf("Current version: %d", current))

	if err = goose.Up(db, "sql"); err != nil {
		log.Fatal("Failed migration: ", err)
	}

	color.Green("Migration ran successfully")
}

func getFilesInDir(dirPath string) ([]string, error) {
	var fileNames []string

	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	// Extract file names from fileInfos
	for _, fileInfo := range fileInfos {
		fileNames = append(fileNames, fileInfo.Name())
	}

	return fileNames, nil
}

func getLatestVersionsFromFiles() (int64, error) {
	fileNames, err := getFilesInDir("internal/migration/sql")
	if err != nil {
		return 0, err
	}

	var latest int64
	for _, fileName := range fileNames {
		version, err := getVersionFromFileName(fileName)
		if err != nil {
			return 0, err
		}
		if version > latest {
			latest = version
		}
	}
	return latest, nil
}

func getVersionFromFileName(fileName string) (int64, error) {
	var version int64
	_, err := fmt.Sscanf(fileName, "%d", &version)
	if err != nil {
		return 0, err
	}
	return version, nil
}
