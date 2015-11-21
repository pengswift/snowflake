package main

import (
	"log"
	"net"

	"github.com/pengswift/internal/guid"
	pb "github.com/pengswift/snowflake/snowflake"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	factory = &guid.GuidFactory{}
)

type server struct{}

func (s *server) GetGUID(ctx context.Context, in *pb.SnowflakeRequest) (*pb.SnowflakeResponse, error) {
	id, err := factory.NewGUID(in.WorkerID)
	return &pb.SnowflakeResponse{Guid: int64(id)}, err
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
