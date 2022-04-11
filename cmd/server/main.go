package main

import (
	"context"
	"log"
	"net"

	"github.com/BernardoAquino/go-grpc-api/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("Mensagem recebida: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grp_Error := grpcServer.Serve(listener)
	if grp_Error != nil {
		log.Fatal(grp_Error)
	}
}
