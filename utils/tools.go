package utils

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func ErrorDebug(err error, message ...string) error {

	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	if err != nil {
		log.Printf("[DEBUG]: %s, error : %s",
			yellow(strings.Join(message, " ")),
			red(err.Error()))
		return err
	}
	return nil
}

func Debug(messages ...string) {
	green := color.New(color.FgGreen).SprintFunc()
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("\n[DEBUG]: Called from [%s], line [%d], func: [%v]\n%s\n",
			file, line, runtime.FuncForPC(pc).Name(),
			green(strings.Join(messages, ", ")),
		)
	}
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
