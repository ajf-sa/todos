package providers

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() *sql.DB {
	dbname := os.Getenv("DBNAME")
	dbpass := os.Getenv("DBPASS")
	dbuser := os.Getenv("DBUSER")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpass, dbname))
	if err != nil {
		panic(err)
	}
	return db
}
