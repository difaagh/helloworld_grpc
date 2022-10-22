package server

import (
	"fmt"

	"github.com/difaagh/helloworld_proto"
	"golang.org/x/net/context"
)

type Server struct {
	helloworld_proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *helloworld_proto.HelloRequest) (*helloworld_proto.HelloReply, error) {
	fmt.Printf("Received: %v", in.GetName())
	return &helloworld_proto.HelloReply{Message: "Hello " + in.GetName() + "! this is from server"}, nil
}
