# transfer-files-with-grpc-golang

## compile proto file

## install

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## generate

```bash
protoc --proto_path=pkg pkg/proto/*.proto --go_out=. --go-grpc_out=.
```

## run server

```bash
go run cmd/server/main.go
```

## run client

```bash
go run -race cmd/client/main.go ./image/name.png
```
