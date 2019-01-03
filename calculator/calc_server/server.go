package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ymgyt/grpc-go-course/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("invoked Calc RPC call: %v\n", req)
	result := req.Arg1 + req.Arg2
	res := &calcpb.CalcResponse{Result: result}

	return res, nil
}

func (*server) PrimeNumberDecompose(req *calcpb.PrimeNumberDecomposeRequest, stream calcpb.CalcService_PrimeNumberDecomposeServer) error {
	fmt.Printf("invoked PrimeNumberDecompose RPC call: %v\n", req)
	n := req.GetPrimeNumber()
	r := int64(2)
	for n > 1 {
		if n%r == 0 {
			res := &calcpb.PrimeNumberDecomposeResponse{
				Factor: r,
			}
			if err := stream.Send(res); err != nil {
				return err
			}
			n /= r
			time.Sleep(200 * time.Millisecond)
		} else {
			r += 1
		}
	}
	return nil
}

func (*server) Average(stream calcpb.CalcService_AverageServer) error {
	avg := func(ns []int64) float64 {
		result := float64(0)
		for _, n := range ns {
			result += float64(n)
		}
		return result / float64(len(ns))
	}

	numbers := make([]int64, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calcpb.AverageResponse{
				Average: avg(numbers),
			})
		}
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, req.Number)
	}
}

func (*server) FindMaximum(stream calcpb.CalcService_FindMaximumServer) error {
	var max int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		n := req.GetNumber()
		if n > max {
			max = n
		}
		if err := stream.Send(&calcpb.FindMaximumResponse{
			Maximum: max,
		}); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (*server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	fmt.Println("received SquareRoot RPC")
	number := req.GetNumber()
	fmt.Println("squre root received: %v\n", number)
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received negative number: %v", number),
		)
	}
	return &calcpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	// register reflection service.
	reflection.Register(s)

	fmt.Println("running gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
