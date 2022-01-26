package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project1/config"
	"project1/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type ProductResponSuccess struct {
	Status  string
	Message string
	Data    []models.ListProductRespon
}

type ResponSuccess struct {
	Status  string
	Message string
}

type PostErrorBind struct {
	Name int
}

type ResponseFailed struct {
	Status  string
	Message string
}

var (
	mock_data_product = models.Product{
		ID:          1,
		Name:        "Product A",
		Price:       3000,
		Description: "Ini Description",
		Qty:         1,
	}
)

func InsertMockDataProductToDB() error {
	var err error
	if err = config.DB.Save(&mock_data_product).Error; err != nil {
		return err
	}
	return nil
}

func TestInsertProductSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	body, err := json.Marshal(mock_data_product)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	if assert.NoError(t, AddProductController(context)) {
		body := rec.Body.String()
		var product ProductResponSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error marshal")
		}
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "success", product.Status)
		assert.Equal(t, "success insert product", product.Message)
	}
}

func TestInsertProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	t.Run("TestInsertProduct_ErrorBind", func(t *testing.T) {
		body, err := json.Marshal(PostErrorBind{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		if assert.NoError(t, AddProductController(context)) {
			body := rec.Body.String()
			var product ProductResponSuccess
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", product.Status)
			assert.Equal(t, "bad request", product.Message)
		}
	})
	t.Run("TestInsertProduct_ErrorFormatName", func(t *testing.T) {
		mock_data_product.Name = "    Product A"
		body, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		if assert.NoError(t, AddProductController(context)) {
			body := rec.Body.String()
			var product ProductResponSuccess
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", product.Status)
			assert.Equal(t, "name cannot less than 5 characters or invalid format", product.Message)
		}
	})
	t.Run("TestInsertProduct_ProductIsExist", func(t *testing.T) {
		mock_data_product.Name = "Product A"
		InsertMockDataProductToDB()
		body, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		if assert.NoError(t, AddProductController(context)) {
			body := rec.Body.String()
			var product ProductResponSuccess
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", product.Status)
			assert.Equal(t, "product name was used, try another one", product.Message)
		}
	})
	t.Run("TestInsertProduct_ErrorDesc", func(t *testing.T) {
		mock_data_product.Description = ""
		mock_data_product.Name = "Product B"
		body, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		if assert.NoError(t, AddProductController(context)) {
			body := rec.Body.String()
			var product ProductResponSuccess
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", product.Status)
			assert.Equal(t, "description cannot be empty", product.Message)
		}
	})
	t.Run("TestInsertProduct_ErrorDB", func(t *testing.T) {
		mock_data_product.Description = "Ini Description"
		config.DB.Migrator().DropTable(&models.Product{})
		body, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		if assert.NoError(t, AddProductController(context)) {
			body := rec.Body.String()
			var product ProductResponSuccess
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "failed", product.Status)
			assert.Equal(t, "internal server error", product.Message)
		}
	})
}
func TestGetProductListSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	mock_data_product.Name = "Product A"
	InsertMockDataProductToDB()

	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/products")

	if assert.NoError(t, GetAllProductController(context)) {
		body := rec.Body.String()
		var product ProductResponSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get all products", product.Message)
		assert.Equal(t, 1, len(product.Data))
		assert.Equal(t, "Product A", product.Data[0].Name)
	}
}

func TestGetProductListFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/products")
	config.DB.Migrator().DropTable(&models.Product{})
	if assert.NoError(t, GetAllProductController(context)) {
		body := rec.Body.String()
		var product ProductResponSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "failed", product.Status)
		assert.Equal(t, "internal server error", product.Message)
	}
}

func TestProductFilterSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	t.Run("TestFilterName", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/filter/:request", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		contex := e.NewContext(req, res)
		contex.SetPath("/products/filter/:request")
		contex.SetParamNames("request")
		contex.SetParamValues("ascending")
		if assert.NoError(t, ProductFilterController(contex)) {
			var product ProductResponSuccess
			body := res.Body.String()
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error unmarshal")
			}
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success get all products by ascending", product.Message)
			assert.Equal(t, "success", product.Status)
			assert.Equal(t, "Product A", product.Data[0].Name)
		}
	})
	t.Run("TestFilterPrice", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/filter/:request", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		contex := e.NewContext(req, res)
		contex.SetPath("/products/filter/:request")
		contex.SetParamNames("request")
		contex.SetParamValues("lower")
		if assert.NoError(t, ProductFilterController(contex)) {
			var product ProductResponSuccess
			body := res.Body.String()
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error unmarshal")
			}
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success get all products price from lower", product.Message)
			assert.Equal(t, "success", product.Status)
			assert.Equal(t, "Product A", product.Data[0].Name)
		}
	})
	t.Run("TestFilterNewProduct", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/filter/:request", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		contex := e.NewContext(req, res)
		contex.SetPath("/products/filter/:request")
		contex.SetParamNames("request")
		contex.SetParamValues("new")
		if assert.NoError(t, ProductFilterController(contex)) {
			var product ProductResponSuccess
			body := res.Body.String()
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error unmarshal")
			}
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success get all products from latest", product.Message)
			assert.Equal(t, "success", product.Status)
			assert.Equal(t, "Product A", product.Data[0].Name)
		}
	})
	t.Run("TestFilterNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/filter/:request", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		contex := e.NewContext(req, res)
		contex.SetPath("/products/filter/:request")
		contex.SetParamNames("request")
		contex.SetParamValues("asss")
		if assert.NoError(t, ProductFilterController(contex)) {
			var product ProductResponSuccess
			body := res.Body.String()
			err := json.Unmarshal([]byte(body), &product)
			if err != nil {
				assert.Error(t, err, "error unmarshal")
			}
			assert.Equal(t, http.StatusNotFound, res.Code)
			assert.Equal(t, "data is doesn't exist", product.Message)
			assert.Equal(t, "success", product.Status)
		}
	})

}
func TestProductFilterFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	req := httptest.NewRequest(http.MethodGet, "/products/filter/:request", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)
	contex.SetPath("/products/filter/:request")
	contex.SetParamNames("request")
	contex.SetParamValues("ascending")
	config.DB.Migrator().DropTable(&models.Product{})
	if assert.NoError(t, ProductFilterController(contex)) {
		var product ProductResponSuccess
		body := res.Body.String()
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error unmarshal")
		}
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "internal server error", product.Message)
		assert.Equal(t, "failed", product.Status)
	}

}
