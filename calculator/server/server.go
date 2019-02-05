package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "github.com/adarsh-carousell/learn-go/calculator/proto"
)

type server struct{}

func (*server) Sum(ctx context.Context,req *proto.SumRequest) (*proto.SumResponse, error) {
	log.Println("Starting call method")
	sum  :=  req.FirstNumber + req.SecondNumber;
	response := &proto.SumResponse{
		SumResult:sum,
	}
	return response, nil
}

func main(){
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err !=nil {
		log.Fatalf("failed to listen")
	}

	s:= grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
