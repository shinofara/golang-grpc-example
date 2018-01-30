package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/shinofara/golang-grpc-example/proto"
	"google.golang.org/grpc"
)

func main() {
	var id int
	var wait int
	var hostname string
	flag.IntVar(&id, "id", 0, "")
	flag.IntVar(&wait, "wait", 0, "")
	flag.StringVar(&hostname, "h", "127.0.0.1", "")
	flag.Parse()

	//sampleなのでwithInsecure
	conn, err := grpc.Dial(fmt.Sprintf("%s:19003", hostname), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewExampleClient(conn)
	message := &pb.GetDataRequest{
		Id:   int32(id),
		Wait: int32(wait),
	}
	res, err := client.GetData(context.TODO(), message)

	fmt.Printf("result:%#v \n", res)

	errprint(err)
}

func errprint(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
