# go-ocpi-api
Satimoto OCPI hub API using golang

## Development

### Generate proto files
Generates the golang code from proto files
```bash
protoc proto/businessdetail.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/image.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/credential.proto --go_out=plugins=grpc:$GOPATH/src
```

### Run
```bash
go run ./cmd/ocpi
```

## Build

### Run
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/main cmd/ocpi/main.go
```
