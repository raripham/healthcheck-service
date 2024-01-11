package database

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DBDriver      string
	DBSource      string
	ServerAddress string
}

func OpenDB() (*sql.DB, error) {

	// Here you can populate the DB params from Env or Config file or use marshaling to fill in the values
	dbParams := Config{
		DBDriver: "postgres",
		DBSource: "postgres://postgres:changeme@localhost:5433/healthcheck?sslmode=disable",
		// DBSource: os.Getenv("DB_URL"),
	}

	db, err := sql.Open(dbParams.DBDriver, dbParams.DBSource)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	return db, nil
}
