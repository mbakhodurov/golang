package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Создание DSN без использования fmt.Sprintf
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return db
}
