package controllers

import (
	"fmt"
	"net/http"
	respon "project1/api/common"
	"project1/lib/databases"
	"project1/models"
	"regexp"

	"github.com/labstack/echo/v4"
)

func AddProductController(c echo.Context) error {
	product := models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, respon.StatusFailed("bad request"))
	}
	// REGEX
	var pattern string
	var matched bool
	// Check Format Name
	pattern = `^(\w+ ?){4}$`
	regex, _ := regexp.Compile(pattern)
	matched = regex.Match([]byte(product.Name))
	if !matched {
		return c.JSON(http.StatusBadRequest, respon.StatusFailed("name cannot less than 5 characters or invalid format"))
	}
	nameCheck, _ := databases.CheckDatabase("name", product.Name)
	if nameCheck > 0 {
		return c.JSON(http.StatusBadRequest, respon.StatusFailed("product name was used, try another one"))
	}
	// Check Format Description
	pattern = `^\w([a-zA-Z0-9()@:%_\+.~#?&//=\n"'\t\\;<>!*-{}]+ ?)*$`
	matched, _ = regexp.Match(pattern, []byte(product.Description))
	if !matched {
		return c.JSON(http.StatusBadRequest, respon.StatusFailed("description cannot be empty"))
	}
	_, err := databases.Insert(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, respon.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusCreated, respon.StatusSuccess("success insert product"))
}

func GetAllProductController(c echo.Context) error {
	products, err := databases.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, respon.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusOK, respon.StatusSuccessData("success get all products", products))
}

func ProductFilterController(c echo.Context) error {
	reqFilter := c.Param("request")
	fmt.Println("reqFilter", reqFilter)
	var err error
	var products []models.ListProductRespon
	jsonRespon := ""
	switch reqFilter {
	case "descending", "ascending":
		products, err = databases.GetFilterByName(reqFilter)
		jsonRespon = "by " + reqFilter
	case "lower", "upper":
		fmt.Println(reqFilter)
		products, err = databases.GetFilterPrice(reqFilter)
		jsonRespon = "price from " + reqFilter
	case "new":
		products, err = databases.GetFilterNewProduct()
		jsonRespon = "from latest"
	default:
		return c.JSON(http.StatusNotFound, respon.StatusSuccessData("data is doesn't exist", nil))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, respon.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusOK, respon.StatusSuccessData("success get all products "+jsonRespon, products))
}
