package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/shinofara/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	// 実行したい実処理をseverに登録する
	pb.RegisterExampleServer(server, Server{})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("start grpc server 0.0.0.0:19003")
		errprint(server.Serve(listenPort))
	}()

	s := <-stop
	switch s {
	// kill -SIGHUP XXXX
	case syscall.SIGHUP:
		fmt.Println("hungup")

	// kill -SIGINT XXXX or Ctrl+c
	case syscall.SIGINT:
		fmt.Println("Warikomi")

	// kill -SIGTERM XXXX
	case syscall.SIGTERM:
		fmt.Println("force stop")

	// kill -SIGQUIT XXXX
	case syscall.SIGQUIT:
		fmt.Println("stop and core dump")
	default:
		fmt.Println("Unknown signal.")
	}

	fmt.Println("Start GracefulStop")
	server.GracefulStop()

}

func errprint(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct{}

func (_ Server) GetData(c context.Context, r *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	fmt.Printf("Request Data %v\n", *r)
	time.Sleep(30 * time.Second)

	return &pb.GetDataResponse{
		Data: fmt.Sprintf("Finish ID %d", r.Id),
	}, nil
}
