package product

import (
	"Project/Go/midtrans/delivery/common"
	"Project/Go/midtrans/entities"
	"Project/Go/midtrans/repository/product"
	awss3 "Project/Go/midtrans/utils/aws-s3"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProductController struct {
	repo product.Product
	// awsConn configs.AppConfig
	conn *session.Session
}

func New(repository product.Product, aws *session.Session) *ProductController {
	return &ProductController{
		repo: repository,
		conn: aws,
	}
}

func (pc *ProductController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		newProductRequest := CreateProductRequest{}
		if err := c.Bind(&newProductRequest); err != nil {
			log.Info(err)
		}
		res, err := pc.repo.Insert(entities.Product{Name: newProductRequest.Name, Code: newProductRequest.Code})

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.StatusBadRequest("Bad request bro"))
		}

		return c.JSON(http.StatusCreated, common.StatusGood(http.StatusAccepted, "Success create Product", res))
	}
}

func (pc *ProductController) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		// conn := awss3.InitS3(pc.awsConn.S3_KEY, pc.awsConn.S3_SECRET, pc.awsConn.S3_REGION)
		productUplaod := UploadImage{}
		c.Bind(&productUplaod)

		file, err := c.FormFile("file")

		if err != nil {
			log.Info(err)
		}

		// res, err := file.Open()
		// if err != nil {
		// 	log.Info("Upload File error : ", err)
		// }
		link := awss3.DoUpload(pc.conn, *file)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"link": link,
		})
	}
}
