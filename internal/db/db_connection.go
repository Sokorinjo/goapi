package db

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	//Simple way of defining DSN:
	// dsn := "vlada:darkcidator@tcp(127.0.0.1:3306)/go_api_db?parseTime=true"

	//Capture connection properties
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:" + os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")

	//Create DSN string from config:
	dsn := cfg.FormatDSN()
	//Get a database handle
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database: " + os.Getenv("DB_NAME"))
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}
