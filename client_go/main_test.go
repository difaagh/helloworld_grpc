package main

import (
	"context"
	"helloworld_grpc/server"
	"log"
	"net"
	"testing"

	"github.com/difaagh/helloworld_proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	s := grpc.NewServer()

	helloworld_proto.RegisterGreeterServer(s, &server.Server{})

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer()), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := helloworld_proto.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &helloworld_proto.HelloRequest{Name: "Test"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	assert.Equal(t, "Hello Test! this is from server", resp.GetMessage())
}
