package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/ymgyt/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)
	doErrorUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	res, err := c.Calc(context.Background(), &calcpb.CalcRequest{
		Arg1: 3,
		Arg2: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("calc rpc call %d + %d = %d\n", 3, 10, res.Result)
}

func doServerStreaming(c calcpb.CalcServiceClient) {
	req := &calcpb.PrimeNumberDecomposeRequest{
		PrimeNumber: 12000000,
	}

	stream, err := c.PrimeNumberDecompose(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PrimeNumberDecomponse Call: %d\n", req.PrimeNumber)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d ", res.Factor)
	}
	fmt.Println()
}

func doClientStreaming(c calcpb.CalcServiceClient) {
	requests := []*calcpb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, req := range requests {
		fmt.Printf("send: %v\n", req)
		stream.Send(req)
		time.Sleep(300 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Client Streaming receive: %v\n", res)
}

func doBiDiStreaming(c calcpb.CalcServiceClient) {
	numbers := []int64{
		10, 100, 1000, 1, 20, 30, 50,
	}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for _, n := range numbers {
			if err := stream.Send(&calcpb.FindMaximumRequest{
				Number: n,
			}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("maximux: %d\n", res.GetMaximum())
	}
}

func doErrorUnary(c calcpb.CalcServiceClient) {
	fmt.Println("starting to do a SquareRoot Unary RPC...")
	number := int32(-10)
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{Number: number})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("we probably sent a negative number !")
				return
			}
		} else {
			log.Fatalf("big error calling SquareRoot: %v", err)
		}
	}
	fmt.Printf("result of square root of %v: %v\n", number, res.GetNumberRoot())
}
