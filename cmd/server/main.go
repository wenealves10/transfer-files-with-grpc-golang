package main

import (
	"log"
	"net"

	"github.com/wenealves10/transfer-files-with-grpc-golang/internal/storage"
	"github.com/wenealves10/transfer-files-with-grpc-golang/internal/upload"
	"github.com/wenealves10/transfer-files-with-grpc-golang/pkg/pb"

	"google.golang.org/grpc"
)

func main() {
	// Initialise TCP listener.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	// Bootstrap upload server.
	uplSrv := upload.NewServer(storage.New("tmp/"))

	// Bootstrap gRPC server.
	rpcSrv := grpc.NewServer()

	// Register and start gRPC server.
	pb.RegisterUploadServiceServer(rpcSrv, uplSrv)
	log.Fatal(rpcSrv.Serve(lis))
}
