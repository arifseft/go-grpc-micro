package repository

import (
	"errors"

	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
	"github.com/stretchr/testify/mock"
)

type StockDecreaseLogRepositoryMock struct {
    Mock mock.Mock
}

func (m *StockDecreaseLogRepositoryMock) FindByOrderId(orderId int64) (*models.StockDecreaseLog, error) {
    var stockDecreaseLog models.StockDecreaseLog

    args := m.Mock.Called(orderId)

    if args.Get(0) == nil {
        return &stockDecreaseLog, errors.New("record not found")
    }

    stockDecreaseLog = args.Get(0).(models.StockDecreaseLog)

    return &stockDecreaseLog, nil
}

func (m *StockDecreaseLogRepositoryMock) CreateStockDecreaseLog(stockDecreaseLog *models.StockDecreaseLog) (*models.StockDecreaseLog, error) {

    args := m.Mock.Called(stockDecreaseLog)

    if args.Get(0) == nil {
        return nil, errors.New("unexpected error")
    }

    return stockDecreaseLog, nil
}
