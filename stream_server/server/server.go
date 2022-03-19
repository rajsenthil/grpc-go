package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	fibonocci "home.com/udemy/grpc/fibonocci/pb"
)

type server struct {
	fibonocci.FibonocciServiceServer
}

func (*server) Fibonocci(req *fibonocci.FibonocciRequest, stream fibonocci.FibonocciService_FibonocciServer) error {
	log.Printf("***Fibonocci streaming started***")

	for i := 0; i < 10; i++ {
		res := &fibonocci.FibonocciResponse{Result: 1}
		stream.Send(res)
		time.Sleep(1000 * 1)
	}

	return nil
}

func main() {
	fmt.Printf("Starting Fibonocci series server...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	s := grpc.NewServer()
	fibonocci.RegisterFibonocciServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Server failed to listen: %v", err)
	}

}
