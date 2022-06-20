package services

import (
	"context"
	"net/http"

	"github.com/arifseft/go-grpc-micro/product-svc/pkg/db"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/pb"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/repository"
)

type Server struct {
    H db.Handler
    ProductRepo repository.IProductRepository
    StockDecreaseLogRepo repository.IStockDecreaseLogRepository
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
    var product models.Product

    product.Name = req.Name
    product.Stock = req.Stock
    product.Price = req.Price

    result, err := s.ProductRepo.CreateProduct(&product)

    if err != nil {
        return &pb.CreateProductResponse{
            Status: http.StatusConflict,
            Error: err.Error(),
        }, err
    }

    return &pb.CreateProductResponse{
        Status: http.StatusCreated,
        Id: result.Id,
    }, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
    product, err := s.ProductRepo.FindOne(req.Id)

    if err != nil {
        return &pb.FindOneResponse{
            Status: http.StatusNotFound,
            Error: err.Error(),
        }, err
    }

    return &pb.FindOneResponse{
        Status: http.StatusOK,
        Data: &pb.FindOneData{
            Id: product.Id,
            Name: product.Name,
            Stock: product.Stock,
            Price: product.Price,
        },
    }, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
    product, err := s.ProductRepo.FindOne(req.Id)

    if err != nil {
        return &pb.DecreaseStockResponse{
            Status: http.StatusNotFound,
            Error: err.Error(),
        }, err
    }

    if product.Stock <= 0 {
        return &pb.DecreaseStockResponse{
            Status: http.StatusConflict,
            Error: "Stock too low",
        }, nil
    }

    log, _ := s.StockDecreaseLogRepo.FindByOrderId(req.OrderId)

    if log.Id != 0 {
        return &pb.DecreaseStockResponse{
            Status: http.StatusConflict,
            Error: "Stock already decreased",
        }, nil
    }

    stock := product.Stock - 1
    _, _ = s.ProductRepo.UpdateProductStock(product, stock)

    log.OrderId = req.OrderId
    log.ProductRefer = product.Id
    _, _ = s.StockDecreaseLogRepo.CreateStockDecreaseLog(log)

    return &pb.DecreaseStockResponse{
        Status: http.StatusOK,
    }, nil
}
