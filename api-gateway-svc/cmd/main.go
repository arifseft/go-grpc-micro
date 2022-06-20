package main

import (
	"log"

	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/auth"
	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/config"
	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/order"
	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    r := gin.Default()

    authSvc := *auth.RegisterRoutes(r, &c)
    product.RegisterRoutes(r, &c, &authSvc)
    order.RegisterRoutes(r, &c, &authSvc)

    r.Run(c.Port)
}
