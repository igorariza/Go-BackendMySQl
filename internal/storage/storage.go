package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var bdMysqlCN = ConnectToDB()

//ConnectToDB func connect generic comment
func ConnectToDB() *sql.DB {
	dbName := os.Getenv("DB_NAME")
	server := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	if server == "" {
		server = "localhost"
	}
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, server, dbPort, dbName)
	db, err := sql.Open("mysql", connURL)

	if err != nil {
		log.Fatalf("Failed to connect to DB via %s: %v", connURL, err)
	}
	if err = db.Ping(); err != nil {
		// log.Fatalf("Failed to ping DB via %s: %v", connURL, err.Error())
		log.Println("Failed to ping a BD" + err.Error())
	}
	log.Println("Connected to DB")
	return db
}

//ChequeoConnection comment generic
func ChequeoConnection() int {
	err := bdMysqlCN.Ping()
	if err != nil {
		return 0
	}
	return 1
}
