package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jessica-millenia/go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		res += fmt.Sprintf("Hello %s %s!\n", req.FirstName, req.LastName)
	}
}
