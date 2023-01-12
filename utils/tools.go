package utils

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
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

const QiniuUrl = "http://img.51cloudclass.com"

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

func GetFileExt(filename string) string {
	pos := strings.LastIndex(filename, ".")
	if pos == -1 {
		return "unknown"
	}
	return filename[pos+1:]
}

func HashString(content []byte) string {
	alg := md5.New()
	alg.Write(content)

	return fmt.Sprintf("%x", alg.Sum(nil))
}

func FullLink(rel string) string {
	return fmt.Sprintf("%s/%s", QiniuUrl, rel)
}

func QiniuUpload(fileHeader *multipart.FileHeader) (string, error) {

	conf := InitConfig()
	accessKey := conf.Qiniu.AccessKey
	secretKey := conf.Qiniu.SecretKey

	bucket := "51cloudclass"
	baseDir := "avatar"
	ext := GetFileExt(fileHeader.Filename)

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

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
		return "", err
	}
	data, err = io.ReadAll(file)
	if err != nil {
		return "", err
	}

	remotePath := fmt.Sprintf("%s/%s.%s",
		baseDir,
		HashString(data),
		ext)

	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(),
		&ret,
		upToken,
		remotePath,
		bytes.NewReader(data),
		dataLen,
		&putExtra)
	if err != nil {
		ErrorDebug(err)
		return "", err
	}
	return FullLink(remotePath), nil
}

type jsonResult map[string]*json.RawMessage

func ReadMockJson(filename string) (jsonResult, error) {
	path, _ := os.Getwd()
	Debug("current path", path)
	if _, err := os.Stat(filename); err != nil {
		return nil, ErrorDebug(err, filename, " is not exists")
	}
	jsonContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, ErrorDebug(err, "cannot read file", filename)
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(jsonContent, &objmap)
	return objmap, err

}
