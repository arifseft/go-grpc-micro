package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/pb"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var stockDecreaseLogRepository = &repository.StockDecreaseLogRepositoryMock{Mock: mock.Mock{}}
var productService = Server{
    ProductRepo: productRepository,
    StockDecreaseLogRepo: stockDecreaseLogRepository,
}

func TestProductService_FindOne_NotFound(t *testing.T) {
    productRepository.Mock.On("FindOne", int64(1)).Return(nil, errors.New("record not found")).Once()

    result, err := productService.FindOne(context.Background(), &pb.FindOneRequest{Id: 1})

    expected := &pb.FindOneResponse{
        Status: http.StatusNotFound,
        Error: err.Error(),
    }

    assert.Equal(t, expected, result)
    assert.NotNil(t, err)
}

func TestProductService_FindOne_Success(t *testing.T) {
    product := models.Product{
        Id: 1,
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    productRepository.Mock.On("FindOne", int64(1)).Return(product, nil).Once()

    result, err := productService.FindOne(context.Background(), &pb.FindOneRequest{Id: 1})

    expected := &pb.FindOneResponse{
        Status: http.StatusOK,
        Data: result.Data,
    }

    assert.Equal(t, expected, result)
    assert.Nil(t, err)
}

func TestProductService_CreateProduct_Success(t *testing.T) {
    product := models.Product{
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    productRepository.Mock.On("CreateProduct", &product).Return(product, nil).Once()

    req := pb.CreateProductRequest{
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    result, err := productService.CreateProduct(context.Background(), &req)

    expected := &pb.CreateProductResponse{
        Status: http.StatusCreated,
        Id: result.Id,
    }

    assert.Equal(t, expected, result)
    assert.Nil(t, err)
}

func TestProductService_CreateProduct_Fail(t *testing.T) {
    product := models.Product{
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    productRepository.Mock.On("CreateProduct", &product).Return(nil, errors.New("error")).Once()

    req := pb.CreateProductRequest{
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    result, err := productService.CreateProduct(context.Background(), &req)

    expected := &pb.CreateProductResponse{
        Status: http.StatusConflict,
        Error: "unexpected error",
    }

    assert.Equal(t, expected, result)
    assert.NotNil(t, err)
}

func TestProductService_DecreaseStock_ProductNotFound(t *testing.T) {
    productRepository.Mock.On("FindOne", int64(1)).Return(nil, errors.New("record not found")).Once()

    req := pb.DecreaseStockRequest{Id: 1, OrderId: 1}

    result, err := productService.DecreaseStock(context.Background(), &req)

    expected := &pb.DecreaseStockResponse{
        Status: http.StatusNotFound,
        Error: err.Error(),
    }

    assert.Equal(t, expected, result)
    assert.NotNil(t, err)
}

func TestProductService_DecreaseStock_ProductOutOfStock(t *testing.T) {
    product := models.Product{
        Id: 1,
        Name: "Product B",
        Stock: 0,
        Price: 100,
    }

    productRepository.Mock.On("FindOne", int64(1)).Return(product, nil).Once()

    req := pb.DecreaseStockRequest{Id: 1, OrderId: 1}

    result, _ := productService.DecreaseStock(context.Background(), &req)

    expected := &pb.DecreaseStockResponse{
        Status: http.StatusConflict,
        Error: "Stock too low",
    }

    assert.Equal(t, expected, result)
}

func TestProductService_DecreaseStock_AlreadyDecreased(t *testing.T) {
    product := models.Product{
        Id: 1,
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    productRepository.Mock.On("FindOne", int64(1)).Return(product, nil).Once()

    stockDecreaseLog := models.StockDecreaseLog{
        Id: 1,
        OrderId: 1,
    }

    stockDecreaseLogRepository.Mock.On("FindByOrderId", int64(1)).Return(stockDecreaseLog, nil).Once()

    req := pb.DecreaseStockRequest{Id: 1, OrderId: 1}

    result, _ := productService.DecreaseStock(context.Background(), &req)

    expected := &pb.DecreaseStockResponse{
        Status: http.StatusConflict,
        Error: "Stock already decreased",
    }

    assert.Equal(t, expected, result)
}

func TestProductService_DecreaseStock_Create(t *testing.T) {
    product := models.Product{
        Id: 1,
        Name: "Product B",
        Stock: 15,
        Price: 100,
    }

    productRepository.Mock.On("FindOne", int64(1)).Return(product, nil).Once()

    stockDecreaseLogRepository.Mock.On("FindByOrderId", int64(1)).Return(nil, nil).Once()

    productRepository.Mock.On("UpdateProductStock", &product, int64(14)).Return(product, nil).Once()

    stockDecreaseLog := models.StockDecreaseLog{
        OrderId: int64(1),
        ProductRefer: 1,
    }

    stockDecreaseLogRepository.Mock.On("CreateStockDecreaseLog", &stockDecreaseLog).Return(stockDecreaseLog, nil).Once()

    req := pb.DecreaseStockRequest{Id: 1, OrderId: 1}

    result, _ := productService.DecreaseStock(context.Background(), &req)

    expected := &pb.DecreaseStockResponse{
        Status: http.StatusOK,
    }

    assert.Equal(t, expected, result)
}
