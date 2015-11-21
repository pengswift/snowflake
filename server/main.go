package main

import (
	"log"

	"github.com/pengswift/internal/guid"
	pb "github.com/pengswift/snowflake/snowflake"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetGUID(ctx context.Context, in *pb.SnowflakeRequest) (*pb.SnowflakeResponse, error) {
	factory := &guidFactory{}
	id, err := factory.NewGUID(in.workerID)

	return &pb.SnowflakeResponse{Guid: id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSnowflakeServiceServer(s, &server{})
	s.Serve(lis)
}
