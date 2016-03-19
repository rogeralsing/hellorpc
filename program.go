package main

import (
	"flag"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/glog"
	"github.com/rogeralsing/hellorpc/myrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	myrpc.RegisterPeopleHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)

	http.ListenAndServe(":8080", mux)
}

func runRpc() {
	lis, _ := net.Listen("tcp", ":9090")
	s := grpc.NewServer()
	myrpc.RegisterPeopleServer(s, &myrpc.Server{})
	s.Serve(lis)
}
//curl  -H 'Content-Type: application/json' -X POST -d '{"name":"abc","age":123}' localhost:8080/api/v1/people
func main() {
	flag.Parse()
	defer glog.Flush()
	go run()
	runRpc()
}
