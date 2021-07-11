package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cota-eng/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {}


func (*server) Greet(ctx context.Context,req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet func invoked %v",req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res,nil
}


func main() {
	fmt.Println("Hello")

	// tcpでgrpc標準のポートへ接続
	lis,err:=net.Listen("tcp","0.0.0.0:50051")
	if err!=nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s:=grpc.NewServer()

	// TODO:要確認
	greetpb.RegisterGreetServiceServer(s,&server{})

	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to serve: %v",err)
	}
}