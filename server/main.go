package main

import (
	"JSMPJ_grpc/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	fmt.Println("This is Server")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum1(), request.GetNum2()
	result := a + b
	return &proto.Response{Result: result}, nil
}

// func (s *server) Sub(ctx context.Context, request *proto.Request) (*proto.Response, error) {
// 	a, b := request.GetNum1(), request.GetNum2()
// 	result := a - b
// 	return &proto.Response{Result: result}, nil
// }

func (s *server) Mul(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum1(), request.GetNum2()
	result := a * b
	return &proto.Response{Result: result}, nil
}

func (s *server) Div(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum1(), request.GetNum2()
	result := a / b
	return &proto.Response{Result: result}, nil
}
