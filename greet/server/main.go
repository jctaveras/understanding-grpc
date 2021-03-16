package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/jctaveras/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := in.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	response := greetpb.GreetResponse{ Result: result }

	return &response, nil
}

func (*server) GreetManyTimes(in *greetpb.GreetManyTimesRequest, stream greetpb.GreetServices_GreetManyTimesServer) error {
	firstName := in.Greeting.FirstName

	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: fmt.Sprintf("Hello %v, number: %v", firstName, i),
		}

		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetServices_LongGreetServer) error {
	result := ""

	for {
		if msg, err := stream.Recv(); err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		} else if err != nil {
			log.Fatalf("Error while reading client stream: %v", msg)
		} else {
			result += fmt.Sprintf("Hello %s! ", msg.Greeting.FirstName)
		}
	}
}

func (*server) GreetEveryone(stream greetpb.GreetServices_GreetEveryoneServer) error {
	for {
		if req, err := stream.Recv(); err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Something when wrong: %v", err)
			
			return err
		} else {
			sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
				Result: fmt.Sprintf("Hello %v !", req.Greeting.FirstName),
			})

			if sendErr != nil {
				log.Fatalf("Something went wrong sending data to the client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (*server) GreetWithDeadLine(ctx context.Context, req *greetpb.GreetWithDeadLineRequest) (*greetpb.GreetWithDeadLineResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client Canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &greetpb.GreetWithDeadLineResponse{
		Result: fmt.Sprintf("Hello %s 👋🏾", req.Greeting.FirstName),
	}, nil
}

func main() {
	fmt.Println("Hello World!")

	list, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	crtFilePath := "ssl/server.crt"
	keyFilePath := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(crtFilePath, keyFilePath)

	if sslErr != nil {
		log.Fatalf("Error loading Credentials: %v", sslErr)
		return
	}

	s := grpc.NewServer(grpc.Creds(creds))
	
	greetpb.RegisterGreetServicesServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
