package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Database() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASS")
	var host = os.Getenv("DB_HOST")
	var dbName = os.Getenv("DB_NAME")
	var credentials = fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True", user, pass, host, dbName)

	database, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Println(err.Error())
	}

	return database
}
