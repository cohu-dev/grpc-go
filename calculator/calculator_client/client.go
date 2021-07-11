package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cota-eng/grpc-go/calculator/calculator_pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect:%v", err)
	}
	defer conn.Close()
	c := calculator_pb.NewCalculatorServiceClient(conn)
	doUnary(c)
}

func doUnary(c calculator_pb.CalculatorServiceClient) {
	req := &calculator_pb.CalculatorRequest{
		NumberOne: 3,
		NumberTwo: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling RPC:%v", err)
	}
	fmt.Printf("Result:%v", res.Result)
}
