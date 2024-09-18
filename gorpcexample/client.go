package main

import (
	"fmt"
	"go-service-call/gorpcexample/common"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var resp common.HelloResponse
	err = client.Call("HelloService.Hello", common.HelloRequest{Name: "world"}, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
