package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/mehulsuthar-000/serverPush/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port         = ":50051"
	noOfMessages = 10
)

// server is used to implement MoneyTransactionServer
type server struct {
	pb.UnimplementedMoneyTransactionServer
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transer...")
	log.Printf("Amount: $%f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)

	// Send noOfMessages to client
	for i := 0; i < noOfMessages; i++ {
		time.Sleep(time.Second * 2)
		// Once task is done, send the successful message
		// bhack to the client
		if err := stream.Send(&pb.TransactionResponse{
			Status:      "good",
			Step:        int32(i),
			Description: fmt.Sprintf("Performing step %d", int32(i)),
		}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transferred amount $%f from %s to %s", in.Amount, in.From, in.To)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
