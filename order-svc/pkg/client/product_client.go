package client

import (
	"context"
	"fmt"

	"github.com/arifseft/go-grpc-micro/order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type IProductServiceClient interface {
    FindOne(req *pb.FindOneRequest) (*pb.FindOneResponse, error)
    DecreaseStock(req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error)
}

type ProductServiceClient struct {
    Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
    cc, err := grpc.Dial(url, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    c := ProductServiceClient{
        Client: pb.NewProductServiceClient(cc),
    }

    return c
}

func (c *ProductServiceClient) FindOne(req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
    return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
    return c.Client.DecreaseStock(context.Background(), req)
}
