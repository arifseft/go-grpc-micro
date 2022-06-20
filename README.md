## go-grpc-micro

### System Diagram
```
graph TD
    A[Client] -->|HTTP| B{{API Gateway}}
    B -->|gRPC| C[Authentication Service]
    B -->|gRPC| D[Product Service]
    B -->|gRPC| E[Order Service]
    E -->|gRPC| D
    C --> F[(DB Auth)]
    D --> G[(DB Product)]
    E --> H[(DB Order)]
```

### Run & Stop Docker Compose
>Run all services & database with docker-compose:
```bash
docker-compose up -d --build
```
>Stop all services & database with docker-compose:
```bash
docker-compose down
```

### Testing
>Run Go unit test.
```bash
cd ./{service_path}
go test -v ./pkg/services
```

### Postman Collection
https://www.getpostman.com/collections/1dc29e94f7d5ec761806
