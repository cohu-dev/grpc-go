package main

import (
	"fmt"
	"log"

	"github.com/cota-eng/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Hello Client")

	// serverとのコネクションを張る localhostのためinsecure
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("could not connect: %v",err)
	}

	// 上位の関数がreturnで終了するまで発火を遅らせる、引数は即時に評価される。
	defer conn.Close()

	c:=greetpb.NewGreetServiceClient(conn)
	fmt.Printf("created client: %v",c)
	
}