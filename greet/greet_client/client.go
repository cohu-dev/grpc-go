package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/cota-eng/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func init() {
	log.SetPrefix("SERVER_STREAMING\n")
}
func main() {
	fmt.Println("Hello Client")

	// serverとのコネクションを張る localhostのためinsecure
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// 上位の関数がreturnで終了するまで発火を遅らせる、引数は即時に評価される。
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)
	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("start Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Yamada",
			LastName:  "Taro",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet API:%v", err)
	}

	log.Printf("response from greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Server",
			LastName:  "Stream",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling server streaming:%v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// if reached end of the file
			break
		}
		if err != nil {
			log.Fatalf("error while read streaming:%v", err)
		}
		log.Printf("Res from GreetManyTimes Message :%v", msg)
	}

}
