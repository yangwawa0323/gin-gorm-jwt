package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func isNotSetOrEmpty(value string) bool {
	if strings.Compare("", value) == 0 {
		return true
	}
	return false
}

func GetConnectURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	if isNotSetOrEmpty(dbuser) || isNotSetOrEmpty(dbpass) ||
		isNotSetOrEmpty(dbhost) || isNotSetOrEmpty(dbport) ||
		isNotSetOrEmpty(dbname) {
		log.Fatal("Please set .env file to correct.")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname)
}
