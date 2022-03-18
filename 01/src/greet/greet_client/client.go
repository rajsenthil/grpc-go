package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"home.com/udemy/grpc/src/greet/greetpb"
)

func main() {
	fmt.Println("Hello World from client.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while dialing from the client: %v", err)
	}

	defer cc.Close()
	//c := greetpb.GreetServiceClient(cc)
	c := greetpb.NewGreetServiceClient(cc)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Senthil",
			LastName:  "Rajendran",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in the response: %v", err)
	}
	log.Printf("Response: %v", res)
}
