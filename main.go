package main

import (
	"database/sql"
	"fmt"
	"log"
	"mini-project-practice-build-rest-api/database"
	"mini-project-practice-build-rest-api/routers"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "liatdibawahlaptop"
	dbname   = "mini-project-practice-build-rest-api"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	var PORT = ":8080"

	psqlInfo := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run(PORT)
	fmt.Println("Successfully Connected")
}
