package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	// PostgreSQL connection string
	connectionString := "postgres://postgres:Passw0rd@localhost:5432/berkicau?sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}
	defer db.Close()

	// Attempt to ping the database to check connectivity
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging the database: %w", err)
	}

	return db, nil
}
