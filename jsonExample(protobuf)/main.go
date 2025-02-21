package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/MehulSuthar-000/protobufs/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	proto_body, _ := proto.Marshal(p)
	p1 := &pb.Person{}
	proto.Unmarshal(proto_body, p1)
	json_body, _ := json.MarshalIndent(p, "", "  ")
	// body is the byte array to be sent over the network
	fmt.Println(proto_body)
	fmt.Println(p1)
	fmt.Println(string(json_body))
}
