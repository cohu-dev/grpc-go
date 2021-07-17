package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cota-eng/grpc-go/calculator/calculator_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculator_pb.CalculatorRequest) (*calculator_pb.CalculatorResponse, error) {
	fmt.Printf("Received:%v", req)
	number1 := req.NumberOne
	number2 := req.NumberTwo
	result := number1 + number2
	res := &calculator_pb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func(*server)PrimeNumberDecomposition(req *calculator_pb.PrimeNumberDecompositionRequest,stream calculator_pb.CalculatorService_PrimeNumberDecompositionServer) error{
	fmt.Printf("Received:%v",req)
	input:=req.GetInput()
	var i int32
	for i= 2;input>1;{
		if input % i == 0 {
			res := &calculator_pb.PrimeNumberDecompositionResponse{
				Result:i,
			}
			input = input / i
			stream.Send(res)
		}else{
			i++
			fmt.Printf("division increment:%v",i)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculator_pb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}
}
