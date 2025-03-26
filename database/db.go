package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sohibjon7731/nectar/config"
)

var DB *sql.DB

func DBConnect() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}

	dsn := config.GetDBDSN()
	fmt.Println("Connecting to:", dsn)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil

}

func ApplyMigrations(DB *sql.DB) {
	migrationFiles := []string{
		"migrations/001_create_users_table.sql",
		"migrations/002_create_categories_table.sql",
		"migrations/003_create_products_table.sql",
	}

	for _, file := range migrationFiles {
		query, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", file, err)
		}
		_, err = DB.Exec(string(query))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file, err)
		}
		fmt.Printf("Migration %s applied successfully!\n", file)
	}
}
