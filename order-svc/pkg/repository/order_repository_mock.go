package repository

import (
	"errors"

	"github.com/arifseft/go-grpc-micro/order-svc/pkg/models"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
    Mock mock.Mock
}

func (m *OrderRepositoryMock) CreateOrder(order *models.Order) (*models.Order, error) {
    args := m.Mock.Called(order)

    if args.Get(0) == nil {
        return nil, errors.New("unexpected error")
    }
    order.Id = 1

    return order, nil
}

func (m *OrderRepositoryMock) DeleteOrder(orderId int64) error {
    args := m.Mock.Called(orderId)

    if args.Get(0) == nil {
        return errors.New("unexpected error")
    }

    return nil
}
