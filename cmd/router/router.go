package router

import (
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}

}

func buildProductsRoutes(r *gin.RouterGroup) {
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

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}
