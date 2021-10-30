package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"calculator/protocol"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I am a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := protocol.NewCalculatorServiceClient(cc)
	Sum(c, 1, 2)
	Sum(c, 6, 8)
	PrimeDecomposition(c, 13)
	PrimeDecomposition(c, 12)
	Average(c, []int64{int64(1), int64(4), int64(5)})
	Average(c, []int64{int64(2), int64(7)})
	Max(c, []int64{int64(2), int64(7), int64(3), int64(9), int64(4), int64(11)})
}

func Sum(c protocol.CalculatorServiceClient, left int64, right int64) {
	fmt.Println("Sum RPC...")
	req := &protocol.SumRequest{Left: left, Right: right}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error requesting sum: %v", err)
	}
	fmt.Printf("Server response was: %d \n", res.GetResult())
}

func PrimeDecomposition(c protocol.CalculatorServiceClient, number int64) {
	fmt.Println("Prime Decomposition RPC...")
	req := &protocol.PrimeDecompositionRequest{Number: number}
	resStream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error requesting prime decomposition: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break;
		}
		if err != nil {
			log.Fatalf("Error while reading prime decomposition reply: %v", err)
		}
		fmt.Printf("Received server response: %d\n", msg.GetDivisor())
	}
}

func Average(c protocol.CalculatorServiceClient, numbers []int64) {
	fmt.Println("Average RPC...")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error requesting average: %v", err)
	}
	for _, n := range numbers {
		err := stream.Send(&protocol.AverageRequest{Number: n})
		if err != nil {
			log.Fatalf("Error sending average sample point: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving average: %v", err)
	}
	fmt.Printf("Received server response: %f\n", res.GetAverage())
}

func Max(c protocol.CalculatorServiceClient, numbers []int64) {
	fmt.Println("Average RPC...")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error requesting max: %v", err)
	}
	for _, n := range numbers {
		err := stream.Send(&protocol.MaxRequest{Number: n})
		if err != nil {
			log.Fatalf("Error sending max sample point: %v", err)
		}
		res, recvErr := stream.Recv()
		if recvErr != nil {
			log.Fatalf("Error receiving latest max: %v", recvErr)
		}
		fmt.Printf("Received server response: %d\n", res.GetMax())
	}
}