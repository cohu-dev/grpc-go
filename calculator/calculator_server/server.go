package main

import (
	"context"

	"github.com/cota-eng/grpc-go/calculator/calculator_pb"
)


type server struct{}

func (*server) Sum(ctx context.Context, req *calculator_pb.CalculatorRequest)(*calculator_pb.CalculatorResponse,error){
	number1 := req.TwoNumber.GetNumber1()
	number2 := req.TwoNumber.GetNumber2()
	result := number1+number2
	res := &calculator_pb.CalculatorResponse{
		Result:result,
	}
	return res,nil
}