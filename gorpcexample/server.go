package main

import (
	"go-service-call/gorpcexample/common"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(req common.HelloRequest, resp *common.HelloResponse) error {
	resp.Message = "Hello " + req.Name
	return nil
}

func main() {
	svc := new(HelloService)
	rpc.Register(svc)
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	log.Printf("server listening at %v", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
