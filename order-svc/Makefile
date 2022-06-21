proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go

test:
	go test -v ./pkg/services

cover:
	go test -v ./pkg/services -coverprofile=/tmp/coverage.out && go tool cover -html=/tmp/coverage.out
