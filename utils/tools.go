package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ErrorDebug(err error) error {
	if err != nil {
		log.Printf("[DEBUG]: %s", err.Error())
		return err
	}
	return nil
}

func Debug(message string) {
	log.Printf("[DEBUG]: %s", message)
}

func LoadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	return nil
}

func RequireAudit() bool {
	var audit bool = false // By default
	LoadDotEnv()
	audit, _ = strconv.ParseBool(os.Getenv("AUDIT_LOG"))
	return audit
}
