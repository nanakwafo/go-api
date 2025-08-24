package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	//create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
