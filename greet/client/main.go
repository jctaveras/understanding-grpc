package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jctaveras/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello From the client")
	crtFilePath := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(crtFilePath, "")

	if sslErr != nil {
		log.Fatalf("Failed to load Certificate: %v", sslErr)
		return
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Error to connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServicesClient(conn)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDirecctionalStreaming(c)

	doUnaryWithDeadline(c, 5 * time.Second)
	doUnaryWithDeadline(c, 1 * time.Second)
}

func doUnary(client greetpb.GreetServicesClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jean Carlos",
			LastName: "Taveras",
		},
	}
	if res, err := client.Greet(context.Background(), req); err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	} else {
		log.Printf("Response from Greet: %v", res.Result)
	}
}

func doServerStreaming(client greetpb.GreetServicesClient) {
	stream, err := client.GreetManyTimes(context.Background(), &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jean Carlos",
			LastName: "Taveras",
		},
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
			log.Printf("Response: %v \n", msg.Result)
		}
	}
}

func doClientStreaming(client greetpb.GreetServicesClient) {
	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Jean Carlos",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Yani",
			},
		},
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	if res, err := stream.CloseAndRecv(); err != nil {
		log.Fatalf("Error receiving the data: %v", err)
	} else {
		fmt.Printf("Recieved Data: %v\n", res)
	}
}

func doBiDirecctionalStreaming(client greetpb.GreetServicesClient) {
	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		
		return
	}

	waitc := make(chan struct{})

	go func () {
		requests := []*greetpb.GreetEveryoneRequest{
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Jean Carlos",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Yani",
				},
			},
		}
		
		for _, req := range requests {
			stream.Send(req)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
	
			if err == io.EOF {
				break
			}
		
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)

				break
			}
	
			fmt.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()
	
	<- waitc
}

func doUnaryWithDeadline(client greetpb.GreetServicesClient, timeout time.Duration) {
	req := &greetpb.GreetWithDeadLineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jean Carlos",
			LastName: "Taveras",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	res, err := client.GreetWithDeadLine(ctx, req)

	if err != nil {
		statusErr, ok := status.FromError(err)

		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit, Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected Error: %v \n", statusErr.Message())
			}
		} else {
			log.Fatalf("Error while calling GreetWithDeadline RPC: %v\n", err)
		}

		return
	}

	fmt.Printf("Message Received: %v\n", res.Result)
}
