package main

import (
	"context"
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

	doUnary(c)
	
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("start Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting:&greetpb.Greeting{
			FirstName:"Yamada",
			LastName:"Taro",
		},
	}
	
	res,err:=c.Greet(context.Background(),req)
	if err!=nil{
		log.Fatalf("error while calling greet API:%v",err)
	}
	
	log.Printf("response from greet: %v",res.Result)
	//  from greet.pb.go 
	// 	type GreetRequest struct {
	// 	state         protoimpl.MessageState
	// 	sizeCache     protoimpl.SizeCache
	// 	unknownFields protoimpl.UnknownFields
	
	// 	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	// }
	
	// type Greeting struct {
	// state         protoimpl.MessageState
	// sizeCache     protoimpl.SizeCache
	// unknownFields protoimpl.UnknownFields
	
	// FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// }
}