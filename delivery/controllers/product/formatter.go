package product

type CreateProductRequest struct {
	Name string `json:"name" form:"name"`
	Code int    `json:"code" form:"code"`
}

type UploadImage struct {
	Name string `form:"name"`
}

type CreateProductResponse struct {
	ID   uint
	Name string
	Code int
}
