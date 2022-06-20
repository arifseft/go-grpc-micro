package order

import (
	"fmt"

	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/config"
	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
    cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewOrderServiceClient(cc)
}
