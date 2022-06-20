package auth

import (
    "fmt"

    "github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/auth/pb"
    "github.com/arifseft/go-grpc-micro/api-gateway-svc/pkg/config"
    "google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
    cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewAuthServiceClient(cc)
}
