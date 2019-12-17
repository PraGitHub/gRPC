package main

import (
	"../api"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	var conn *grpc.ClientConn
	port := 8080

	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewPingClient(conn)
	for {
		clientName := fmt.Sprintf("gRPC_client_%s", GetRandString(25))
		response, err := c.SayHello(context.Background(), &api.PingMessage{Sender: clientName, Message: fmt.Sprintf("Hi, this is %s", clientName)})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from %s server: %s", response.Sender, response.Message)
		time.Sleep(5 * time.Second)
	}
}
