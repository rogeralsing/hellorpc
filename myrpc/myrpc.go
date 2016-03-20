package myrpc

import (
	"golang.org/x/net/context"
)

type Server struct{}

func (this Server) CreatePerson(context context.Context, request *CreatePersonRequest) (*CreatePersonResponse, error) {
	println("Fick en RPC Request...", request.Name, request.Age)
	response := &CreatePersonResponse{
		Result: "Det gck bra",
	}
	return response, nil
}

func (this Server) GetPerson(context context.Context, request *GetPersonRequest) (*GetPersonResponse, error) {
	//println("Fick en RPC Request...", request.Name)
	response := &GetPersonResponse{
		Name: "Roger",
		Age:  40,
	}
	return response, nil
}

func (this Server) GetPeople(empty *EmptyMessage, stream People_GetPeopleServer) error {
	response := &GetPersonResponse{
		Name: "Roger",
		Age:  40,
	}
	for i:=0;i<1000000 ;i++  {
		stream.Send(response)
	}
	return nil
}
