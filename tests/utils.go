package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	r := gin.Default()
	
	var prodList []products.Product
	prodList = append(prodList, products.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	prodList = append(prodList, products.Product{
		ID:          "mock2",
		SellerID:    "FEX112AC2",
		Description: "generic product2",
		Price:       123.55,
	})
	repo := products.NewRepository(prodList)
	service := products.NewService(repo)
	handler := products.NewHandler(service)

	prodRoute := r.Group("/api/v1/products")
	prodRoute.GET("", handler.GetProducts)
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
