package repository

import (
	"errors"

	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
    Mock mock.Mock
}

func (m *ProductRepositoryMock) FindOne(id int64) (*models.Product, error) {
    var product models.Product

    args := m.Mock.Called(id)

    if args.Get(0) == nil {
        return &product, errors.New("record not found")
    }

    product = args.Get(0).(models.Product)

    return &product, nil
}

func (m *ProductRepositoryMock) UpdateProductStock(product *models.Product, stock int64) (*models.Product, error) {

    args := m.Mock.Called(product, stock)

    if args.Get(0) == nil {
        return nil, errors.New("unexpected error")
    } else if args.Get(1) == nil {
        return nil, errors.New("unexpected error")
    }

    return product, nil
}

func (m *ProductRepositoryMock) CreateProduct(product *models.Product) (*models.Product, error) {

    args := m.Mock.Called(product)

    if args.Get(0) == nil {
        return nil, errors.New("unexpected error")
    }

    return product, nil
}
