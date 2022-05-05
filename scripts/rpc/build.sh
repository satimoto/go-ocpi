protoc ocpirpc/businessdetail.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/image.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/credential.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/command.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/session.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/cdr.proto --go_out=plugins=grpc:$GOPATH/src
protoc ocpirpc/token.proto --go_out=plugins=grpc:$GOPATH/src