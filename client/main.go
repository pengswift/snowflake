package main

import (
	"log"

	pb "github.com/pengswift/snowflake/snowflake"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSnowflakeServiceClient(conn)

	r, err := c.GetGUID(context.Background(), &pb.SnowflakeRequest{WorkerID: 0})
	if err != nil {
		log.Fatalf("could not getuuid: %v", err)
	}
	log.Printf("GUID: %d", r.Guid)
}
