package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

// Define Unix domain socket to which server is listening to
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid()) // Match server's socket naming
	return s
}

func main() {
	// Connect to the Unix domain socket
	sockname := coordinatorSock()
	client, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// Call the Hello method on the Coordinator
	args := "World"
	var reply map[string]interface{}

	//rpc call to the server function
	err = client.Call("Coordinator.Hello", args, &reply)
	if err != nil {
		fmt.Println("RPC call error:", err)
		return
	}

	// Print the response
	fmt.Println("Response from server:", reply)
}
