package databaserc

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LouisDecaudaveine/rekord_cloud/internal/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	utils.Check(err)
}

func InitDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL not set in .env file")
	}

	db, err := sql.Open("postgres", connStr)
	utils.Check(err)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
