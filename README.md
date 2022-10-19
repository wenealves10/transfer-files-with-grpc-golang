# transfer-files-with-grpc-golang

## compile proto file

```bash
protoc --proto_path=pkg pkg/proto/*.proto --go_out=.
```

```bash
protoc --go_out=plugins=grpc:. --go_opt=paths=. pkg/proto/*.proto
```