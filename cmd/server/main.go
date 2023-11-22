package main

import (
	"fmt"

	"github.com/majidehamide/gorest-api2/internal/database"
)

func Run() error {
	fmt.Println("Application run without error")

	db, err := database.NewDatabase()

	if err != nil {
		fmt.Println("Could not connect to database : %w", err)
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database : %w", err)
		return err
	}
	fmt.Println("Successfully connect to database")
	return nil
}
func main() {
	fmt.Println("start the application")
	if err := Run(); err != nil {
		fmt.Println("Error when start application")
	}
}
