package configs

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port      int16
	DB        string
	DB_NAME   string
	DB_PORT   int16
	HOST      string
	USERNAME  string
	PASSWORD  string
	S3_KEY    string
	S3_SECRET string
	S3_REGION string
}

var synchronizer = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	synchronizer.Lock()
	defer synchronizer.Unlock()
	if appConfig == nil {
		appConfig = InitConfig()
	}

	return appConfig
}

func InitConfig() *AppConfig {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Info(err)
	}
	var defaultConfig AppConfig = AppConfig{8000, "mysql", "local_db", 3306, "localhost", "root", "", "", "", ""}

	res, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Warn(err)
	}

	defaultConfig.Port = int16(res)
	defaultConfig.DB = os.Getenv("DB")
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	res, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Warn(err)
	}

	defaultConfig.DB_PORT = int16(res)
	defaultConfig.HOST = os.Getenv("HOST")
	defaultConfig.USERNAME = os.Getenv("USERNAME")
	defaultConfig.PASSWORD = os.Getenv("PASSWORD")
	defaultConfig.S3_KEY = os.Getenv("S3-KEY")
	defaultConfig.S3_SECRET = os.Getenv("S3-SECRET")
	defaultConfig.S3_REGION = os.Getenv("S3-REGION")

	log.Info(defaultConfig)

	return &defaultConfig
}
