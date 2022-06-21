package client

import (
	"errors"

	"github.com/arifseft/go-grpc-micro/order-svc/pkg/pb"
	"github.com/stretchr/testify/mock"
)

type ProductServiceClientMock struct {
    Mock mock.Mock
}

func (m *ProductServiceClientMock) FindOne(req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
    var res *pb.FindOneResponse

    args := m.Mock.Called(req)

    if args.Get(0) == nil {
        return res, errors.New("record not found")
    }

    res = args.Get(0).(*pb.FindOneResponse)

    return res, nil
}

func (m *ProductServiceClientMock) DecreaseStock(req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
    var res *pb.DecreaseStockResponse

    args := m.Mock.Called(req)

    if args.Get(0) == nil {
        return nil, errors.New("unexpected error")
    }

    res = args.Get(0).(*pb.DecreaseStockResponse)

    return res, nil
}
