package main

import (
	"../api"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// main start a gRPC server and waits for connection
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	port := 8080
	// create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening to port ", port)

	// create a server instance
	serverName := fmt.Sprintf("gRPC_server_%s", GetRandString(25))
	log.Println("Creating server instance with the identifier ", serverName)
	s := api.Server{Name: serverName}
	log.Println("Done creating server instance with the identifier ", serverName)

	// create a gRPC server object
	log.Println("Creating gRPC server object")
	grpcServer := grpc.NewServer()
	log.Println("Done creating gRPC server object")

	// attach the Ping service to the server
	log.Println("Attching the ping servie to the server")
	api.RegisterPingServer(grpcServer, &s)
	log.Println("Done attaching the ping service to the server")

	// start the server
	log.Println("Starting the gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("Done starting the gRPC server")
}
