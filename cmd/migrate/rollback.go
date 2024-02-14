package main

import (
	"bufio"
	"database/sql"
	"fmt"
	migrations "github.com/MrRytis/apple-business-manager/internal/migration"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	color.Green("Running migration up")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	color.White("Enter your migration version to down to (integer): ")

	// Read the user's input
	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
		return
	}

	// Remove spaces and new lines from the input
	input = strings.ReplaceAll(input, " ", "_")
	input = strings.ReplaceAll(input, "\n", "")

	version, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal("Failed to parse version: ", err)
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

	if err = goose.DownTo(db, "sql", version); err != nil {
		log.Fatal("Failed migration: ", err)
	}

	color.Green("Migration ran successfully")

}
