package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/shinofara/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 0, "")
	flag.Parse()

	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewExampleClient(conn)
	message := &pb.GetDataRequest{
		Id: int32(id),
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
