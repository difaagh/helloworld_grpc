package main

import (
	"fmt"
	"helloworld_grpc/server"
	"net"

	"github.com/difaagh/helloworld_proto"
	"google.golang.org/grpc"
)

const port = 8080

func main() {
	serv, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	helloworld_proto.RegisterGreeterServer(s, &server.Server{})
	fmt.Printf("server Listening at %v\n", serv.Addr())
	if err := s.Serve(serv); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
