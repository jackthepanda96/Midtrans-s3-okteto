package product

import (
	"Project/Go/midtrans/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) *ProductRepository {
	return &ProductRepository{db: dbConn}
}

func (pr *ProductRepository) Get() ([]entities.Product, error) {
	resArr := []entities.Product{}

	if err := pr.db.Find(&resArr).Error; err != nil {
		return nil, err
	} else if len(resArr) == 0 {
		return nil, gorm.ErrEmptySlice
	}

	return resArr, nil
}

func (pr *ProductRepository) Insert(newProduct entities.Product) (entities.Product, error) {
	if err := pr.db.Create(&newProduct).Error; err != nil {
		log.Info(err)
		return newProduct, err
	}

	return newProduct, nil
}
