package services

import (
	"context"
	"net/http"

	"github.com/arifseft/go-grpc-micro/order-svc/pkg/client"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/db"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/models"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/pb"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/repository"
)

type Server struct {
    H db.Handler
    OrderRepo repository.IOrderRepository
    ProductSvc client.IProductServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    product, err := s.ProductSvc.FindOne(&pb.FindOneRequest{
        Id: req.ProductId,
    })

    if err != nil {
        return &pb.CreateOrderResponse{
            Status: http.StatusBadRequest,
            Error: err.Error(),
        }, nil
    } else if product.Status >= http.StatusNotFound {
        return &pb.CreateOrderResponse{
            Status: product.Status,
            Error: product.Error,
        }, nil
    } else if product.Data.Stock < req.Quantity {
        return &pb.CreateOrderResponse{
            Status: http.StatusConflict,
            Error: "Stock too less",
        }, nil
    }

    order := models.Order{
        Price: product.Data.Price,
        ProductId: product.Data.Id,
        UserId: req.UserId,
    }

    result, err := s.OrderRepo.CreateOrder(&order)

    if err != nil {
        return &pb.CreateOrderResponse{
            Status: http.StatusConflict,
            Error: err.Error(),
        }, err
    }

    res, err := s.ProductSvc.DecreaseStock(&pb.DecreaseStockRequest{
        Id: req.ProductId,
        OrderId: order.Id,
    })

    if err != nil {
        return &pb.CreateOrderResponse{
            Status: http.StatusBadRequest,
            Error: err.Error(),
        }, nil
    } else if res.Status == http.StatusConflict {
        s.H.DB.Delete(&models.Order{}, order.Id)

        return &pb.CreateOrderResponse{
            Status: http.StatusConflict,
            Error: res.Error,
        }, nil
    }

    return &pb.CreateOrderResponse{
        Status: http.StatusCreated,
        Id: result.Id,
    }, nil
}
