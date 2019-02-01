package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "github.com/adarsh-carousell/learn-go/calculator/proto"
)

type server struct{}

func (*server) Sum(ctx context.Context,req *proto.SumRequest) (*proto.SumResponse, error) {
	sum  :=  req.FirstNumber + req.SecondNumber;
	response := &proto.SumResponse{
		SumResult:sum,
	}
	return response, nil
}

func main(){
	_, err := net.Listen("tcp", "0.0.0.0.:50051")
	if err !=nil {
		log.Fatalf("failed to listen")
	}

	s:= grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &server{})
}
