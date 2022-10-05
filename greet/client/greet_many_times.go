package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jessica-millenia/go-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Jessica",
		LastName:  "Millenia",
	})

	if err != nil {
		log.Fatalf("Could not GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
