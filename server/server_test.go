package server

import (
	"context"
	"testing"

	"github.com/difaagh/helloworld_proto/stub/go/protos/helloworld"
	"github.com/stretchr/testify/assert"
)

var is_called = false

type mockServer struct {
	*Server
}

func (*mockServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	is_called = true
	return nil, nil
}

func newServer() *mockServer {
	return &mockServer{&Server{}}
}

func TestServer(t *testing.T) {
	s := newServer()
	s.SayHello(context.Background(), &helloworld.HelloRequest{Name: "Test"})
	assert.True(t, is_called)
}
