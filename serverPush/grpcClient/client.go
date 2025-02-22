package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mehulsuthar-000/serverPush/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}

	// Listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// If there are no more messages, break the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Front-End or Android App
	request := &pb.TransactionRequest{
		Amount: 100,
		From:   "1234",
		To:     "5678",
	}

	// Contact the server and print out its response.
	ReceiveStream(client, request)
}
