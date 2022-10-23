package server

import (
	"fmt"

	"github.com/difaagh/helloworld_proto/stub/go/protos/helloworld"
	"golang.org/x/net/context"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName() + "! this is from server"}, nil
}
