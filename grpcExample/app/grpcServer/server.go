package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/mehulsuthar-000/grpcExample/protofiles"
)

const (
	port = ":50051"
)

// server is used to create MoneyTransactionServer.
type server struct {
	pb.UnimplementedMoneyTransactionServer
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Got request for money Transfer....")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	// Do database logic here....
	return &pb.TransactionResponse{Confirmation: true}, nil
}

// register the server
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v",
			err)
	}
	s := grpc.NewServer()

	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",
			err)
	}
}
