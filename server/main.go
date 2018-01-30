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

	pb "github.com/shinofara/golang-grpc-example/proto"
	"google.golang.org/grpc"
)

var status string

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
		status = "server is running"
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

	status = "server is start graceful stop"
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

	var wait int32 = 0
	if w := r.Wait; w > 0 {
		wait = w
	}
	time.Sleep(time.Duration(wait) * time.Second)

	return &pb.GetDataResponse{
		Data: fmt.Sprintf("Finish ID %d, Server status is %s", r.Id, status),
	}, nil
}
