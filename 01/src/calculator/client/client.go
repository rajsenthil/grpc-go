package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	calc "home.com/udemy/grpc/src/calculator/pb"
)

func main() {
	fmt.Println("Calculator sum of two numbers.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while dialing from the client: %v", err)
	}

	defer cc.Close()
	c := calc.NewCalculatorClient(cc)
	req := &calc.CalcRequest{
		Calc: &calc.Calc{
			FirstNumber:  12,
			SecondNumber: 13,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in the client response: %v", err)
	}
	log.Printf("Sum of two numbers: %v", res)
}
