package main

import (
	"context"
	"log"

	pb "github.com/jessica-millenia/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  10,
		SecondNumber: 13,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum result: %d\n", res.Result)
}
