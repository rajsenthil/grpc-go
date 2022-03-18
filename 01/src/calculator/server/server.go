package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	calc "home.com/udemy/grpc/src/calculator/pb"
)

type server struct {
	calc.CalculatorServer
}

func (*server) Sum(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {
	firstNumber := req.GetCalc().GetFirstNumber()
	secondNumber := req.GetCalc().GetSecondNumber()
	sum := firstNumber + secondNumber
	log.Printf("First number: %v, Second number: %v and their sum = %v", firstNumber, secondNumber, sum)
	return &calc.CalcResponse{Sum: firstNumber + secondNumber}, nil
}

func main() {
	fmt.Println("Sum Calculator")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calc.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server calculator: %v", err)
	}
}
