package main

import (
	"context"
	"fmt"
	"log"

	"github.com/k-washi/example-golang-gRPC/unary/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello. I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "tanaka",
			LastName:  "tarou",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("error while calling greet rpc: %v", err)
	}

	fmt.Printf("Response from greet server %v", res)
}
