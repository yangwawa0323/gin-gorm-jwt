package utils

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func ErrorDebug(err error, message ...string) error {

	red := color.New(color.FgHiRed).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	if err != nil {
		log.Printf("[DEBUG]: %s, error : %s",
			yellow(strings.Join(message, " ")),
			red(err.Error()))
		return err
	}
	return nil
}

func Debug(messages ...string) {
	green := color.New(color.FgHiGreen).SprintFunc()
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

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func QiniuUpload(fileHeader *multipart.FileHeader) error {

	conf := InitConfig()
	accessKey := conf.Qiniu.AccessKey
	secretKey := conf.Qiniu.SecretKey

	bucket := "51cloudclass"

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	Debug(upToken)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "upload.ico",
		},
	}

	var data []byte
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	data, err = io.ReadAll(file)
	if err != nil {
		return err
	}

	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, upToken, accessKey, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		ErrorDebug(err)
		return err
	}
	Debug("Qiniu upload result:", ret.Key, ret.Hash)
	return nil
}
