FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest
RUN apk add --no-cache make
COPY --from=builder /app/main .
EXPOSE 50053
ENTRYPOINT ["./main"]
