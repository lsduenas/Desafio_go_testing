package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

var (
	ErrRepositoryProductNotFound = errors.New("repository: product not found")
)

type repository struct {
	Data []Product
}

func NewRepository(data []Product) Repository {
	return &repository{Data: data}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	// Fill database
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	prodList = append(prodList, Product{
		ID:          "mock2",
		SellerID:    "FEX112AC2",
		Description: "generic product2",
		Price:       123.55,
	})
	// validate seller ID in database
	var prodListBySellerID []Product
	var flag bool
	for _, value := range prodList {
		if value.SellerID == sellerID {
			prodListBySellerID = append(prodListBySellerID, value)
			flag = true
		}
	}

	if flag {
		return prodListBySellerID, nil
	} else {
		return []Product{}, ErrRepositoryProductNotFound
	}

}
