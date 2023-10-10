package products

type MockRepository struct {
	DataOnGetAll []Product
	ErrOnGetAll error
	// calls -> spy
}

func NewMockRepository (dataOnGetAll []Product, errOnGetAll error) *MockRepository{
	return &MockRepository{DataOnGetAll: dataOnGetAll, ErrOnGetAll: errOnGetAll}
}

// return products data 
func (m *MockRepository) GetAllBySellerMock(sellerID string) (*[]Product, error) {
	return &m.DataOnGetAll, m.ErrOnGetAll
}