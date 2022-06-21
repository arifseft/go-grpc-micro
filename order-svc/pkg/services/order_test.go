package services

import (
	"context"
	"net/http"
	"testing"

	"github.com/arifseft/go-grpc-micro/order-svc/pkg/client"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/models"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/pb"
	"github.com/arifseft/go-grpc-micro/order-svc/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productServiceClient = &client.ProductServiceClientMock{Mock: mock.Mock{}}
var orderRepository = &repository.OrderRepositoryMock{Mock: mock.Mock{}}
var orderService = Server{
    ProductSvc: productServiceClient,
    OrderRepo: orderRepository,
}

func TestOrderService_CreateOrder_Success(t *testing.T) {
    product := pb.FindOneResponse{
        Status: http.StatusOK,
        Data: &pb.FindOneData{
            Id: 1,
            Name: "Product B",
            Stock: 15,
            Price: 150,
        },
    }

    productServiceClient.Mock.On("FindOne", &pb.FindOneRequest{Id: int64(1)}).Return(&product, nil).Once()

    req := pb.CreateOrderRequest{
        ProductId: 1,
        Quantity: 2,
        UserId: 3,
    }

    orderRequest := models.Order{
        Price: 150,
        ProductId: 1,
        UserId: 3,
    }

    orderResponse := models.Order{
        Id: 1,
        Price: 150,
        ProductId: 1,
        UserId: 3,
    }

    orderRepository.Mock.On("CreateOrder", &orderRequest).Return(&orderResponse, nil).Once()

    decreaseStockRequest := &pb.DecreaseStockRequest{
        Id: req.ProductId,
        OrderId: orderResponse.Id,
    }

    decreaseStockResponse := &pb.DecreaseStockResponse{
        Status: http.StatusOK,
    }

    productServiceClient.Mock.On("DecreaseStock", decreaseStockRequest).Return(decreaseStockResponse, nil).Once()

    result, err := orderService.CreateOrder(context.Background(), &req)

    expected := &pb.CreateOrderResponse{
        Status: http.StatusCreated,
        Id: orderResponse.Id,
    }

    assert.Equal(t, expected, result)
    assert.Nil(t, err)
}
