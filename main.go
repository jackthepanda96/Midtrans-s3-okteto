package main

import (
	"Project/Go/midtrans/configs"
	productCont "Project/Go/midtrans/delivery/controllers/product"
	"Project/Go/midtrans/delivery/routes"
	"Project/Go/midtrans/repository/product"
	awss3 "Project/Go/midtrans/utils/aws-s3"
	"Project/Go/midtrans/utils/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	config := configs.GetConfig()

	db := mysql.ConnDB(config)
	mysql.Migrate(db)

	repo := product.New(db)
	awsConn := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)
	prodCont := productCont.New(repo, awsConn)

	routes.RegisterPath(e, prodCont)

	log.Fatal(e.Start(":8000"))

	// c := midtranspay.InitConnection()
	// midtranspay.CreateTransaction(c)
}
