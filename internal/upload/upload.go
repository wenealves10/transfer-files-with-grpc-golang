package upload

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/wenealves10/transfer-files-with-grpc-golang/pkg/pb"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.UploadServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: pb.NewUploadServiceClient(conn),
	}
}

func (c Client) Upload(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(12000*time.Second))
	defer cancel()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	fil, err := os.Open(file)
	if err != nil {
		return "", err
	}

	// Maximum buffer size is 2MB.
	buf := make([]byte, 2*1024*1024)

	infoFile, err := fil.Stat()
	if err != nil {
		return "", err
	}

	sizeFile := infoFile.Size()

	var total int64

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if err := stream.Send(&pb.UploadRequest{Chunk: buf[:num]}); err != nil {
			return "", err
		}

		total += int64(num)

		fmt.Printf("Progress: %.2f%% \r", float64(total)/float64(sizeFile)*100)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}
