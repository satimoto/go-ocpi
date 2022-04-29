protoc proto/businessdetail.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/image.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/credential.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/command.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/session.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/cdr.proto --go_out=plugins=grpc:$GOPATH/src
protoc proto/token.proto --go_out=plugins=grpc:$GOPATH/src