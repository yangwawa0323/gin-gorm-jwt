package audit

import (
	"errors"
	"log"
	"os"
	"time"

	myerr "github.com/yangwawa0323/gin-gorm-jwt/utils/errors"
	"gorm.io/gorm"
)

var Errors = myerr.Errors

type Storage interface {
	Save(*AuditLog) error
}

type FileStorage struct {
	Filename string
}

type GormStorage struct {
	DB *gorm.DB
}

func (db *GormStorage) Save(al *AuditLog) error {
	return errors.New(Errors[myerr.NotImplemented])
}

func (fs *FileStorage) Save(al *AuditLog) error {
	file, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		panic(Errors[myerr.OpenFileError])
	}
	defer file.Close()

	logMsg := "[LOG] "
	logMsg += time.Now().Format("2016-01-02 15:04:05")
	logMsg += " : "
	logMsg += al.Content + "\r"
	_, err = file.WriteString(logMsg)
	if err != nil {
		log.Fatal(err)
	}
	return errors.New(Errors[myerr.NotImplemented])
}

type AuditLog struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}

func (al *AuditLog) Log(storage Storage) error {
	return storage.Save(al)
}

func Log(message string) {
	audit := AuditLog{
		Content: message,
	}
	fs := &FileStorage{
		Filename: "audit.log",
	}
	audit.Log(fs)
}
