package routes

import (
	"Project/Go/midtrans/delivery/controllers/product"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, pc *product.ProductController) {
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "host=${host}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("products", pc.Insert())
	e.POST("upload", pc.Upload())
}
