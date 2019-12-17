package main

import (
	"fmt"
	"log"
	"time"

	"../api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var crMessage chan string
var interval int

func clientRoutine(c api.PingClient, id int) {
	for {
		clientName := fmt.Sprintf("gRPC_client_%s", GetRandString(25))
		log.Printf("clientRoutine %d :: Pinging server with identifier %s", id, clientName)
		response, err := c.SayHello(context.Background(), &api.PingMessage{Sender: clientName, Message: fmt.Sprintf("Hi, this is %s", clientName)})
		if err != nil {
			log.Printf("clientRoutine %d :: Error when calling SayHello: %s with identifier %s", id, err, clientName)
		} else {
			log.Printf("clientRoutine %d :: Response from %s server: %s", id, response.Sender, response.Message)
		}
		crMessage <- fmt.Sprintf("clientRoutine %d :: Sleeping for %d seconds", id, interval)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	var conn *grpc.ClientConn
	port := 8080

	interval = 5
	crMessage = make(chan string)

	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewPingClient(conn)

	for i := 0; i <= 4; i++ {
		go clientRoutine(c, i)
	}

	for {
		log.Println(<-crMessage)
	}
}
