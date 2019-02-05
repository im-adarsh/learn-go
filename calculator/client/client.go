package main

import (
	"context"
	"fmt"
	"github.com/adarsh-carousell/learn-go/calculator/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer cc.Close()

	c := proto.NewCalculatorServiceClient(cc)

	callSum(c)
}

func callSum(c proto.CalculatorServiceClient){
	req := &proto.SumRequest{
		FirstNumber:1,
		SecondNumber:2,
	}

	res, err := c.Sum(context.Background(),req)
	if err != nil {
		log.Fatal()
	}

	log.Printf("Response from the greet: %v" , res.GetSumResult())
}