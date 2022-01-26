package routers

import (
	"project1/api/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/products", controllers.AddProductController)
	e.GET("/products", controllers.GetAllProductController)
	e.GET("/products/filter/:request", controllers.ProductFilterController)
	return e
}
