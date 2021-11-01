package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"log"
	"math"

	"calculator/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalculatorServer struct {}

func (s *CalculatorServer) Sum(ctx context.Context, req *protocol.SumRequest) (*protocol.SumResponse, error) {
	if ctx.Err() == context.Canceled {
		fmt.Println("The client canceled the request")
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}
	return &protocol.SumResponse{Result: req.GetLeft() + req.GetRight()}, nil
}

func (s *CalculatorServer) PrimeDecomposition(req *protocol.PrimeDecompositionRequest, stream protocol.CalculatorService_PrimeDecompositionServer) error {
	k := int64(2)
    n := req.GetNumber()
    for n > 1 {
		if n % k == 0 {
			res := &protocol.PrimeDecompositionResponse{Divisor: k}
			if stream.Context().Err() == context.Canceled {
				fmt.Println("The client canceled the request")
				return status.Error(codes.Canceled, "the client canceled the request")
			}
			stream.Send(res)
			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (s *CalculatorServer) Average(stream protocol.CalculatorService_AverageServer) error {
	sum := int64(0)
	count := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(-1)
			if count > 0 {
				average = float64(sum) / float64(count)
			}
			if stream.Context().Err() == context.Canceled {
				fmt.Println("The client canceled the request")
				return status.Error(codes.Canceled, "the client canceled the request")
			}
			return stream.SendAndClose(&protocol.AverageResponse{Average: average})
		}
		if err != nil {
			return err
		}
		count += int64(1)
		sum += req.GetNumber()
	}

	return nil
}

func (s *CalculatorServer) Max(stream protocol.CalculatorService_MaxServer) error {
	max := int64(0)
	first := true
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if first || max < req.GetNumber() {
			max = req.GetNumber()
			if stream.Context().Err() == context.Canceled {
				fmt.Println("The client canceled the request")
				return status.Error(codes.Canceled, "the client canceled the request")
			}
			stream.Send(&protocol.MaxResponse{Max: max})
			first = false
		}
	}
	return nil
}

func (s *CalculatorServer) SquareRoot(ctx context.Context, req *protocol.SquareRootRequest) (*protocol.SquareRootResponse, error) {
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %d", number),
		)
	}

	if ctx.Err() == context.Canceled {
		fmt.Println("The client canceled the request")
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}
	return &protocol.SquareRootResponse{Root: math.Sqrt(float64(number))}, nil
}

func main() {
	fmt.Println("Hello, I am a server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protocol.RegisterCalculatorServiceServer(s, &CalculatorServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}