package main

import (
	"context"
	pb "distribute_store/rpc_definition"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
	"io"
)

var (
	serverAddress = flag.String("server_address", "localhost:8000", "The server address in the format of host:port")
)

func getValue(client pb.DistributeStoreClient, key int32) {
	log.Printf("Retrieving value(s) of Key: %v ", key)
	keyToSend := &pb.Key{Key: key}
	stream, err := client.Get(context.Background(), keyToSend)
	if err != nil {
		log.Fatalf("%v.Get(%v) = _, %v: ", client, keyToSend, err)
	}
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Get(%v) = _, %v: ", client, keyToSend, err)
		}
		log.Printf("The value(s) of the Key %v is: %v\n", key, value.GetValue())
	}
}

func appendValue(client pb.DistributeStoreClient, key int32, value string) {
	log.Printf("Sending Key: %v with Value: %v", key, value)
	message := &pb.Message{Key: &pb.Key{Key: key}, Value: &pb.Value{Value: value}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Append(ctx, message)
	if err != nil {
		log.Fatalf("%v.Append(%v) = _, %v: ", client, message, err)
	}
}

func putValue(client pb.DistributeStoreClient, key int32, value string) {
	log.Printf("Sending Key: %v with Value: %v", key, value)
	message := &pb.Message{Key: &pb.Key{Key: key}, Value: &pb.Value{Value: value}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Put(ctx, message)
	if err != nil {
		log.Fatalf("%v.Put(%v) = _, %v: ", client, message, err)
	}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddress, opts...)
	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDistributeStoreClient(conn)
	putValue(client, 1, "1")
	appendValue(client, 1, "2")
	putValue(client, 2, "1")
	getValue(client, 1)
	getValue(client, 2)
}