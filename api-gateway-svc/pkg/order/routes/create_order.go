package routes

import (
	"context"
	"net/http"

	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
    ProductId int64 `json:"productId"`
    Quantity int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
    body := CreateOrderRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    userId := ctx.GetInt64("userId")

    res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
        ProductId: body.ProductId,
        Quantity: body.Quantity,
        UserId: userId,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(http.StatusCreated, &res)
}
