package database

import (
	"database/sql"
	"fmt"
	"go-template/config"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Connect() error {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.GetEnvConfig("DATABASE_USERNAME"),
		config.GetEnvConfig("DATABASE_PASSWORD"),
		config.GetEnvConfig("DATABASE_HOSTNAME"),
		config.GetEnvConfig("DATABASE_PORT"),
		config.GetEnvConfig("DATABASE_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("connect fail" + err.Error())
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("Connected to database")
	return nil

}
