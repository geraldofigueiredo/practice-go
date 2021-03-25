package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	log.Info("Server configuration...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("Failed to listen on por 9000: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC server on por 9000: %v", err)
	}
}
