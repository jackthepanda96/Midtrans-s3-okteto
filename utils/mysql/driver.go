package mysql

import (
	"Project/Go/midtrans/configs"
	"Project/Go/midtrans/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB(appConfig *configs.AppConfig) *gorm.DB {
	connString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		appConfig.USERNAME,
		appConfig.PASSWORD,
		appConfig.HOST,
		appConfig.DB_PORT,
		appConfig.DB_NAME,
	)
	log.Info(connString)
	res, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Product{})
}
