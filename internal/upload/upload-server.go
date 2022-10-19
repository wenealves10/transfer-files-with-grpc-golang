package upload

import (
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/wenealves10/transfer-files-with-grpc-golang/internal/storage"
	"github.com/wenealves10/transfer-files-with-grpc-golang/pkg/pb"
)

type Server struct {
	storage storage.Manager
	pb.UnimplementedUploadServiceServer
}

func NewServer(storage storage.Manager) Server {
	return Server{
		storage: storage,
	}
}

func (s Server) Upload(stream pb.UploadService_UploadServer) error {
	name := "some-unique-name.png"
	file := storage.NewFile(name)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := s.storage.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&pb.UploadResponse{Name: name})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}
