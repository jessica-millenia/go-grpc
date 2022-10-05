package main

import (
	"log"
	"net"

	pb "github.com/jessica-millenia/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var address string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", address)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
