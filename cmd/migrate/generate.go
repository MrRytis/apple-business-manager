package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	color.Green("Running migration generate")

	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	color.White("Enter your migration name: ")

	// Read the user's input
	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
		return
	}

	// Remove spaces and new lines from the input
	input = strings.ReplaceAll(input, " ", "_")
	input = strings.ReplaceAll(input, "\n", "")

	directory := "internal/migration/sql"
	fileName := getFileName(input)
	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		color.Red("Error:", err)
		return
	}
	defer file.Close()

	// Write the SQL content to the file
	_, err = file.WriteString(getSqlContent())
	if err != nil {
		color.Red("Error:", err)
		return
	}

	// Flush the file to ensure all data is written
	err = file.Sync()
	if err != nil {
		color.Red("Error:", err)
		return
	}

	color.Green(fmt.Sprintf("Migration \"%s\" generated successfully", fileName))
}

func getFileName(name string) string {
	return fmt.Sprintf("%s_%s.sql", time.Now().Format("200601021504"), name)
}

func getSqlContent() string {
	return `-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
--CREATE TABLE users (
--    id SERIAL PRIMARY KEY,
--    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--    deleted_at TIMESTAMP,
--);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- DROP TABLE users;`
}
