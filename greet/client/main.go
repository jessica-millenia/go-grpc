package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jessica-millenia/go-grpc/greet/proto"
)

var address string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	doGreetManyTimes(c)
}
