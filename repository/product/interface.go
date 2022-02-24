package product

import "Project/Go/midtrans/entities"

type Product interface {
	Get() ([]entities.Product, error)
	Insert(newUser entities.Product) (entities.Product, error)
}
