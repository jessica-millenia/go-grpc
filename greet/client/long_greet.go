package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jessica-millenia/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Jessica", LastName: "Millenia"},
		{FirstName: "Jessica2", LastName: "Millenia2"},
		{FirstName: "Jessica3", LastName: "Millenia3"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Could not LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
