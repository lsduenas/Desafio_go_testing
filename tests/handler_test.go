package tests

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductsBySellerIDHappyPath(t *testing.T) {
	// Arrange
	// crear el Server y definir las Rutas
	r := createServer()
	responseExpected := `
	[{
		"id": "mock2",
		"seller_id": "FEX112AC2",
		"description": "generic product2",
		"price": 123.55
	}]`

	// crear Request del tipo GET y Response para obtener el resultado
	requestExpected, responseObtained := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC2", "")
	// add query param
	//requestExpected.URL.Query().Add("seller_id", "FEX112AC2")

	// Act
	// Ejecutando el handler
	r.ServeHTTP(responseObtained, requestExpected)
	t.Log(requestExpected)
	log.Println("request ", requestExpected)
	log.Println("response OBTAINED", responseObtained)
	// Assert
	assert.Equal(t, http.StatusOK, responseObtained.Code)
	assert.JSONEq(t, responseExpected, responseObtained.Body.String())
}

func TestGetProductBySellerIDSadPath(t *testing.T) {
	// Arrange
	// crear el Server y definir las Rutas
	r := createServer()
	responseExpected := `
	{
		"error": "repository: product not found"
	}`
	// crear Request del tipo GET y Response para obtener el resultado
	requestExpected, responseObtained := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC3", "")

	// add query param
	//requestExpected.URL.Query().Add("seller_id", "FEX112AC2")

	// Act
	// Ejecutando el handler
	r.ServeHTTP(responseObtained, requestExpected)
	log.Println("request ", requestExpected)
	log.Println("response OBTAINED", responseObtained)
	// Assert
	assert.Equal(t, http.StatusNotFound, responseObtained.Code)
	assert.JSONEq(t, responseExpected, responseObtained.Body.String())
}
