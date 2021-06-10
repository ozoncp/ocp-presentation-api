package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ozoncp/ocp-presentation-api/internal/api"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
)

func runGrpc() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterPresentationAPIServer(s, api.NewPresentationAPI())

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()

	go runGrpc()
}
