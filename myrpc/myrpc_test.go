package myrpc

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
	"net"
	"log"
	"golang.org/x/net/context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestStreamClient(t *testing.T) {
	//arrange
	const address = "127.0.0.1:8090"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	server := Server{}
	RegisterPeopleServer(s,&server)
	go s.Serve(lis)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	client := NewPeopleClient(conn)

	//act
	stream,err := client.GetPeople(context.Background(), &EmptyMessage{})

	for i:=0;i<1000000 ;i++  {
		stream.Recv()
		//assert
		//assert.Equal(t,"Roger",response.Name)
	}
	_,err = stream.Recv()
	assert.NotNil(t,err)
}