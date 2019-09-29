package main

import (
	"log"
	"net"

	"github.com/u03013112/ss-config/config"
	pb "github.com/u03013112/ss-pb/config"
	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen %s", port)
	s := grpc.NewServer()
	pb.RegisterSSConfigServer(s, &config.Srv{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
