package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anindyaSeng/microservices/mvc-grpc/api"
	"github.com/anindyaSeng/microservices/mvc-grpc/controllers"
	"github.com/anindyaSeng/microservices/mvc-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	api.UnimplementedUserServiceServer
}

// GetUser implements api.GetUser
func (s *server) GetUser(ctx context.Context, in *api.UserIDRequest) (*api.UserDataReply, error) {
	log.Printf("Received User data request for User Id: %v", in.GetId())

	user, err := services.GetUser(in.GetId())

	if err != nil {
		return &api.UserDataReply{IsFound: false, UserData: nil}, fmt.Errorf("User for user id %v not found", in.GetId())
	}
	return &api.UserDataReply{IsFound: true, UserData: user}, nil

	/*if id := in.GetId(); id != 123 {
		return &api.UserDataReply{IsFound: false, UserData: nil}, fmt.Errorf("User for user id %v not found", id)
	}
	return &api.UserDataReply{IsFound: true,
		UserData: &api.UserData{
			Id:        123,
			FirstName: "ab",
			LastName:  "cd",
			Email:     "abcd@gmail.com",
		},
	}, nil*/

}

// StartApp -- initialize the app and listen for grpc calls
func StartApp() {
	// Start the controller
	controllers.InitUserController()
	// Now listen to gRPC calls
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, &server{})
	reflection.Register(s) // Needed for reflection API, only then grpCui finds it
	//https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
