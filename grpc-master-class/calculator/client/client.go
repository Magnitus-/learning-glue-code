package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"calculator/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
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
	SquareRoot(c, 9)
	SquareRoot(c, -4)
}

func Sum(c protocol.CalculatorServiceClient, left int64, right int64) {
	fmt.Println("Sum RPC...")
	req := &protocol.SumRequest{Left: left, Right: right}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.Sum(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error requesting sum: %v", err)
		}
		return
	}
	fmt.Printf("Server response was: %d \n", res.GetResult())
}

func PrimeDecomposition(c protocol.CalculatorServiceClient, number int64) {
	fmt.Println("Prime Decomposition RPC...")
	req := &protocol.PrimeDecompositionRequest{Number: number}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resStream, err := c.PrimeDecomposition(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error requesting prime decomposition: %v", err)
		}
		return
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break;
		}
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				if statusErr.Code() == codes.DeadlineExceeded {
					fmt.Println("Request timed out")
				} else {
					fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
				}
			} else {
				log.Fatalf("Error while reading prime decomposition reply: %v", err)
			}
			return
		}
		fmt.Printf("Received server response: %d\n", msg.GetDivisor())
	}
}

func Average(c protocol.CalculatorServiceClient, numbers []int64) {
	fmt.Println("Average RPC...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.Average(ctx)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error requesting average: %v", err)
		}
		return
	}
	for _, n := range numbers {
		err := stream.Send(&protocol.AverageRequest{Number: n})
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				if statusErr.Code() == codes.DeadlineExceeded {
					fmt.Println("Request timed out")
				} else {
					fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
				}
			} else {
				log.Fatalf("Error sending average sample point: %v", err)
			}
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error receiving average: %v", err)
		}
		return
	}
	fmt.Printf("Received server response: %f\n", res.GetAverage())
}

func Max(c protocol.CalculatorServiceClient, numbers []int64) {
	fmt.Println("Average RPC...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.Max(ctx)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error requesting max: %v", err)
		}
		return
	}

	stop := make(chan struct {})

	go func() {
		for _, n := range numbers {
			err := stream.Send(&protocol.MaxRequest{Number: n})
			if err != nil {
				statusErr, ok := status.FromError(err)
				if ok {
					if statusErr.Code() == codes.DeadlineExceeded {
						fmt.Println("Request timed out")
					} else {
						fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
					}
				} else {
					log.Fatalf("Error sending max sample point: %v", err)
				}
				stream.CloseSend()
				return
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, recvErr := stream.Recv()
			if recvErr == io.EOF {
				break
			}
			if recvErr != nil {
				statusErr, ok := status.FromError(recvErr)
				if ok {
					if statusErr.Code() == codes.DeadlineExceeded {
						fmt.Println("Request timed out")
					} else {
						fmt.Printf("Unforeseen grpc error: %v\n", statusErr)
					}
				} else {
					log.Fatalf("Error receiving latest max: %v", recvErr)
				}
				close(stop)
				return
			}
			fmt.Printf("Received server response: %d\n", res.GetMax())
		}
		close(stop)
	}()

	<-stop
}

func SquareRoot(c protocol.CalculatorServiceClient, number int64) {
	fmt.Println("Square Root RPC...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.SquareRoot(ctx, &protocol.SquareRootRequest{Number: number})
	
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			//grpc format error
			fmt.Println("Error message: ", resErr.Message())
			fmt.Println("Error Code: ", resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("Square root argument must be greater or equal to 0")
			} else if resErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Request timed out")
			} else {
				fmt.Printf("Uknown grpc error %v\n", resErr)
			}
		} else {
			log.Fatalf("Error getting square root: %v", resErr)
		}
		return
	}

	fmt.Printf("Received server response: %f\n", res.GetRoot())
}