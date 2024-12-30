# INSTALL gRPC

![grpc-logo](https://grpc.io/img/logos/grpc-logo.png)


## install protoc

protoc is Protobuf compiler

```bash
curl -sL https://github.com/protocolbuffers/protobuf/releases/download/v29.2/protoc-29.2-linux-x86_64.zip \
    | busybox unzip -d $$(go env GOPATH) - \
    && chmod +x $$(go env GOPATH)/bin/protoc
```

- [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)


## install protoc-gen-go

protoc-gen-go is Protobuf compiler for golang 

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.1
```

## install protoc-gen-go-grpc

protoc-gen-go-grpc is gRPC compiler for golang 

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
```

## GENERATE CODE

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```

## install runtime

install google.golang.org/protobuf

```bash
go get google.golang.org/protobuf@v1.36.1
```

install google.golang.org/grpc

```bash
go get google.golang.org/grpc@v1.69.2
```
