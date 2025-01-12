package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

// Coordinator structure with an example RPC method.
type Coordinator struct{}

func (c *Coordinator) Hello(args string, reply *map[string]interface{}) error {
	*reply = map[string]interface{}{"response": "hello " + args}
	return nil
}

func (c *Coordinator) server() {
	// Register the Coordinator object for RPC
	rpc.Register(c)

	//Used to set up an HTTP handler for RPC communication.
	rpc.HandleHTTP()

	// Fetch unix domain socket file path
	sockname := coordinatorSock()
	// Remove old socket file if exists
	os.Remove(sockname)

	// Listen on the Unix domain socket
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// Serve RPC requests over HTTP
	go http.Serve(l, nil)
}

// Create a unix domain socket for IPC
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid()) // Append current user ID
	return s
}

func main() {
	c := Coordinator{}
	c.server()

	// Prevent main from exiting
	select {}
}
