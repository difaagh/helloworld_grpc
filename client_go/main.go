package main

import (
	"context"
	"fmt"
	"time"

	"github.com/difaagh/helloworld_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("cannot connect: %v\n", err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := helloworld_proto.NewGreeterClient(conn)
	resp, err := c.SayHello(ctx, &helloworld_proto.HelloRequest{Name: "go client"})
	if err != nil {
		fmt.Printf("cannot SayHello: %v\n", err)
	}
	fmt.Printf("Response: %s\n", resp.GetMessage())
}
