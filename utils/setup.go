package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func isNotSetOrEmpty(value string) bool {
	return strings.Compare("", value) == 0
}

func GetCertFiles(conf *InitYamlConfig) (cert string, key string) {
	cert = conf.Certs.CertFile
	key = conf.Certs.KeyFile
	if !IsExists(cert) || !IsExists(key) {
		panic(Errors[KeyFileNotExists])
	}
	return
}

func IsExists(file string) bool {
	_, err := os.Stat(file)
	return !errors.Is(err, os.ErrNotExist)

}

func GetConnectURI() string {
	// LoadDotEnv()
	config := InitConfig()

	dbuser := config.Database.User
	dbpass := config.Database.Password
	dbhost := config.Database.Host
	dbport := config.Database.Port
	dbname := config.Database.DbName

	if isNotSetOrEmpty(dbuser) || isNotSetOrEmpty(dbpass) ||
		isNotSetOrEmpty(dbhost) || isNotSetOrEmpty(strconv.Itoa(int(dbport))) ||
		isNotSetOrEmpty(dbname) {
		ErrorDebug(errors.New(Errors[PanicConfigFile]))
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname)
}

/**
*  configure struction
 */
type database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port,omitempty"`
	DbName   string `yaml:"dbname"`
}

type cors struct {
	AllowOrigin string `yaml:"allow_origin,omitempty"`
}

type logSet struct {
	Audit bool `yaml:"audit,omitempty"`
}

type mailbox struct {
	AdminEmail    string `yaml:"admin_email"`
	AdminPassword string `yaml:"admin_password"`
	Host          string `yaml:"host"`
	Port          int64  `yaml:"port,omitempty"`
}

type server struct {
	Backend  backend  `yaml:"backend,omitempty"`
	Frontend frontend `yaml:"frontend,omitempty"`
}

type backend struct {
	Url string `yaml:"url"`
}

type frontend struct {
	Url string `yaml:"url"`
}

type certs struct {
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

type InitYamlConfig struct {
	Database database `yaml:"database"`
	CORS     cors     `yaml:"cors,omitempty"`
	Log      logSet   `yaml:"log,omitempty"`
	Mailbox  mailbox  `yaml:"mailbox,omitempty"`
	Server   server   `yaml:"server,omitempty"`
	Certs    certs    `yaml:"certs"`
}

// end of config structure.

func InitConfig() *InitYamlConfig {
	config := &InitYamlConfig{}
	confFile, err := os.ReadFile("conf/app.yaml")
	if err != nil {
		ErrorDebug(err, Errors[FileNotExist])
		panic(fmt.Sprintf("%s :[%s]", Errors[FileNotExist], "conf/app.yaml"))
	}

	if err := yaml.Unmarshal(confFile, config); err != nil {
		ErrorDebug(err, Errors[UnmarshalWrong], "conf/app.yaml")
		panic(Errors[UnmarshalWrong])
	}

	// Debug(fmt.Sprintf("%#v", config))
	return config
	// for k, v := range data {
	// }
}
