package product

import (
	"fmt"

	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/config"
	"github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
    cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewProductServiceClient(cc)
}
