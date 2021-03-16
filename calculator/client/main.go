package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jctaveras/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not Dial due to: %v", err)
	}

	defer conn.Close()

	client := calcpb.NewCalculatorClient(conn)

	// doUnary(client)
	// doServerStreaming(client)
	// doClientStreaming(client)
	// doBiDirectionalStreaming(client)
	doErrorUnary(client)
}

func doUnary(client calcpb.CalculatorClient) {
	req := &calcpb.AddRequest{
		Addition: &calcpb.Addition{
			Params: []int32{3, 10},
		},
	}
	if res, err := client.Add(context.Background(), req); err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	} else {
		log.Printf("Response from Greet: %v", res.Result)
	}
}

func doServerStreaming(client calcpb.CalculatorClient) {
	stream, err := client.PrimaryNumberDecomposition(context.Background(), &calcpb.PNDRequest{
		Number: 536,
	})

	if err != nil {
		log.Fatalf("Error While streaming: %v", err)
	}

	for {
		if msg, err := stream.Recv(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error While streaming: %v", err)
		} else {
			log.Printf("Response: %v \n", msg.Number)
		}
	}
}

func doClientStreaming(client calcpb.CalculatorClient) {
	stream, err := client.ComputeAvg(context.Background());

	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	requests := []*calcpb.ComputeAvgRequest{
		{ Number: 1 },
		{ Number: 2 },
		{ Number: 3 },
		{ Number: 4 },
	}

	for _, req := range requests {
		stream.Send(req)
	}

	if response, err := stream.CloseAndRecv(); err != nil {
		log.Fatalf("Error Reading the average: %v", err)
	} else {
		fmt.Printf("The Avg is: %.2f", response.Result)
	}
}

func doBiDirectionalStreaming(client calcpb.CalculatorClient) {
	stream, err := client.FindMax(context.Background())

	if err != nil {
		log.Fatalf("Something went wrong creating streaming client: %v", err)
		
		return
	}
	
	waitc := make(chan struct{})

	go func() {
		requests := []*calcpb.FindMaxRequest{
			{ Number: 1 },
			{ Number: 5 },
			{ Number: 3 },
			{ Number: 6 },
			{ Number: 20 }, 
		}

		for _, req := range requests {
			fmt.Printf("Sent Value: %d\n", req.Number)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	go func ()  {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("Something went wrong receiving data from server: %v", err)
				close(waitc)
				return
			}

			fmt.Printf("Max Number: %d\n", res.MaxNumber)
		}
	}()

	<-waitc
}

func doErrorUnary(client calcpb.CalculatorClient) {
	number := float64(100)
	
	if res, err := client.Sqrt(context.Background(), &calcpb.SqrtRequest{ Number: number }); err != nil {
		resErr, ok := status.FromError(err)

		if ok {
			fmt.Printf("Message: %v, Status: %v \n", resErr.Message(), resErr.Code())
			
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("A negative value was sent")
			}
		} else {
			log.Fatalf("Something went wrong: %v", err)
		}
	} else {
		fmt.Printf("The Sqrt of %.2f is %.2f\n", number, res.Result)
	}

	if res, err := client.Sqrt(context.Background(), &calcpb.SqrtRequest{ Number: (number * -1) }); err != nil {
		resErr, ok := status.FromError(err)

		if ok {
			fmt.Printf("Message: %v, Status: %v \n", resErr.Message(), resErr.Code())
			
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("A negative value was sent")
			}
		} else {
			log.Fatalf("Something went wrong: %v", err)
		}
	} else {
		fmt.Printf("The Sqrt of %.2f is %.2f\n", number, res.Result)
	}
}
