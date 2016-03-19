package main

import (
	"flag"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/glog"
	"github.com/rogeralsing/hello-rpc/myrpc"
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
func main() {
	flag.Parse()
	defer glog.Flush()
	go run()
	runRpc()
}
