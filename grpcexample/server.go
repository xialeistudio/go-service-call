package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "go-service-call/grpcexample/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// 中间件示例
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startAT := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Printf("method: %s, duration: %v", info.FullMethod, time.Since(startAT))
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
