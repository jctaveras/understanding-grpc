package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/jctaveras/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {}

func maxNumber(a int32, b int32) int32 {
	if a > b {
		return a
	} 

	return b
}

func (*server) Add(ctx context.Context, in *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	params := in.Addition.Params
	ctx = context.WithValue(ctx, "result", int32(0))

	for _, param := range params {
		// Just Use the ctx instead of the easy stuff
		// Rember to cast the value from context.Value() since the function
		// Returns an interface{}
		ctx = context.WithValue(ctx, "result", ctx.Value("result").(int32) + param)
	}

	response := calcpb.AddResponse{ Result: ctx.Value("result").(int32) }

	return &response, nil
}

func (*server) PrimaryNumberDecomposition(in *calcpb.PNDRequest, stream calcpb.Calculator_PrimaryNumberDecompositionServer) error {
	number := in.Number
	prime := int32(2)

	for number > 1 {
		if number % prime == 0 {
			number = number / prime

			stream.Send(&calcpb.PNDResponse{
				Number: prime,
			})
		} else {
			prime++
		}
	}

	return nil
}

func (*server) ComputeAvg(stream calcpb.Calculator_ComputeAvgServer) error {
	counter := float64(0)
	ctx := context.WithValue(context.Background(), "result", int32(0))
	for {
		if number, err := stream.Recv(); err == io.EOF {
			return stream.SendAndClose(&calcpb.ComputeAvgResponse{
				Result: float64(ctx.Value("result").(int32)) / counter,
			})
		} else if err != nil {
			log.Fatalf("Error while reciving data: %v", err)
		} else {
			counter++
			ctx = context.WithValue(ctx, "result", ctx.Value("result").(int32) + number.Number)
		}
	}
}

func (*server) FindMax(stream calcpb.Calculator_FindMaxServer) error {
	current := int32(0)

	for {
		if req, err := stream.Recv(); err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Something went wrong recibing data: %v", err)
	
			return err
		} else {
			current = maxNumber(current, req.Number)
			
			sendErr := stream.Send(&calcpb.FindMaxResponse{
				MaxNumber: current,
			})
	
			if sendErr != nil {
				log.Fatalf("Something went wrong sending data to the client: %v", err)
	
				return sendErr
			}
		}
	}
}

func (*server) Sqrt(ctx context.Context, req *calcpb.SqrtRequest) (*calcpb.SqrtResponse, error) {
	number := req.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &calcpb.SqrtResponse{ Result: math.Sqrt(number) }, nil
}

func main() {
	list, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	calcpb.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
