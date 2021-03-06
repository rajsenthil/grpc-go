package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	fibonocci "home.com/udemy/grpc/fibonocci/pb"
)

func main() {
	fmt.Println("Fibonocci client starting...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dialing to Fibonocci server: %v", err)
	}

	defer cc.Close()
	c := fibonocci.NewFibonocciServiceClient(cc)
	res, err := c.Fibonocci(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Failed to get fibonocci streaming: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error with server streaming: %v", err)
		}

		log.Printf("Streamin: %v", msg)
	}
}
