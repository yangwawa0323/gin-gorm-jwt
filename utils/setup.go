package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func isNotSetOrEmpty(value string) bool {
	return strings.Compare("", value) == 0
}

func GetConnectURI() string {
	LoadDotEnv()

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
