package main

import (
	"context"
	"fmt"
	"io"
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
	// doUnary(c)
	doServerStreaming(c)
}

// func doUnary(c calculator_pb.CalculatorServiceClient) {
// 	req := &calculator_pb.CalculatorRequest{
// 		NumberOne: 3,
// 		NumberTwo: 2,
// 	}
// 	res, err := c.Sum(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error when calling RPC:%v", err)
// 	}
// 	fmt.Printf("Result:%v", res.Result)
// }

func doServerStreaming(c calculator_pb.CalculatorServiceClient){
	req := &calculator_pb.PrimeNumberDecompositionRequest{
		Input:30,
	}
	resStream,err:=c.PrimeNumberDecomposition(context.Background(),req)
	if err!=nil{
		log.Fatalf("Error:%v",err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while read streaming:%v", err)
		}
		// log.Printf("Res from GreetManyTimes Message :%v", res)
		fmt.Printf("RESULT:%v\n",res.GetResult())
	}
}