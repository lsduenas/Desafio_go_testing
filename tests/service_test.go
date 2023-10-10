package tests

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerIDHappyPath(t *testing.T) {
	// Arrange
	// Seller ID Param
	sellerID := "FEX112AC"

	// Products expected from database
	var productsParam []products.Product
	productsParam = append(productsParam, products.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	// Product expected from service
	var productsExpected []products.Product
	productsExpected = append(productsExpected, products.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	// Mock de repository
	mockRepository := products.NewMockRepository(productsParam, nil)

	// Dependencies TO-DO
	repo := products.NewRepository(mockRepository.DataOnGetAll)

	service := products.NewService(repo)

	// Act
	productsObtained, err := service.GetAllBySeller(sellerID)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, productsExpected, productsObtained)
}

func TestGetAllBySellerIDSadPath(t *testing.T) {
	// Arrange
	// Seller ID Param
	sellerID := "FEX112ACXX"

	// Products expected from database
	var productsParam []products.Product

	// Product expected from service
	var productsExpected []products.Product

	// Error message expected
	errExpected := errors.New("repository: product not found")

	// Mock de repository
	mockRepository := products.NewMockRepository(productsParam, errExpected)

	// Dependencies TO-DO
	repo := products.NewRepository(mockRepository.DataOnGetAll)

	service := products.NewService(repo)

	// Act
	productsObtained, err := service.GetAllBySeller(sellerID)

	// Assert
	assert.Equal(t, errExpected.Error(), err.Error())
	assert.Equal(t, productsExpected, productsObtained)
}
