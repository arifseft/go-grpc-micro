package repository

import (
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/db"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/models"
)

type IOrderRepository interface {
    CreateOrder(order *models.Order) (*models.Order, error)
    DeleteOrder(orderId int64) error
}

type OrderRepository struct {
    H db.Handler
}

func (r *OrderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
    result := r.H.DB.Create(&order)

    if result.Error != nil {
        return nil, result.Error
    }

    return order, nil
}

func (r *OrderRepository) DeleteOrder(orderId int64) error {
    result := r.H.DB.Delete(&models.Order{}, orderId)

    if result.Error != nil {
        return result.Error
    }

    return nil
}
