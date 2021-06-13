package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ozoncp/ocp-presentation-api/internal/api"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	server := grpc.NewServer()
	reflection.Register(server)

	desc.RegisterPresentationAPIServer(server, api.NewPresentationAPI())

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	flag.Parse()

	runGrpc()
}
