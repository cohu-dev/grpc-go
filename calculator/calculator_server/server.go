package main

import (
	"context"
	"log"
	"net"

	"github.com/cota-eng/grpc-go/calculator/calculator_pb"
	"google.golang.org/grpc"
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

func main(){
	lis,err:=net.Listen("tcp","0.0.0.0:50051")
	if err!=nil{
		log.Fatalf("Failed to listen: %v",err)
	}
	s:=grpc.NewServer()
	calculator_pb.RegisterCalculatorServiceServer(s,&server{})
	if err:=s.Serve(lis);err!=nil{
		log.Fatalln("failed to serve:",err)
	}
}