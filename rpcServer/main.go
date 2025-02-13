package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill the pointer to the reply with the current Unix timestamp
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// register the TimeServer object as an RPC service
	// and listen for incoming requests

	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	// Listen for incoming requests on the specified port
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error:", e)
	}
	http.Serve(l, nil)
}
