package repository

import (
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/db"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
)

type IStockDecreaseLogRepository interface {
    FindByOrderId(orderId int64) (*models.StockDecreaseLog, error)
    CreateStockDecreaseLog(stockDecreaseLog *models.StockDecreaseLog) (*models.StockDecreaseLog, error)
}

type StockDecreaseLogRepository struct {
    H db.Handler
}

func (r *StockDecreaseLogRepository) FindByOrderId(orderId int64) (*models.StockDecreaseLog, error) {
    var stockDecreaseLog models.StockDecreaseLog

    result := r.H.DB.Where(&models.StockDecreaseLog{OrderId: orderId}).First(&stockDecreaseLog)

    if result.Error != nil {
        return &stockDecreaseLog, result.Error
    }

    return &stockDecreaseLog, nil
}

func (r *StockDecreaseLogRepository) CreateStockDecreaseLog(stockDecreaseLog *models.StockDecreaseLog) (*models.StockDecreaseLog, error) {
    result := r.H.DB.Create(&stockDecreaseLog)

    if result.Error != nil {
        return nil, result.Error
    }

    return stockDecreaseLog, nil
}
