package main

import (
	"os"
	"strconv"
	"context"
	pb "distribute_store/rpc_definition"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	)

var (
	memory = make(map[int32][]string)
)

type keyValueServer struct {
	pb.UnimplementedDistributeStoreServer
}

func(kvs *keyValueServer) Get(key *pb.Key, stream pb.DistributeStore_GetServer) error {
	logOperation("Get")
	keySent := key.Key
	if values, found := memory[keySent]; found {
		for _, value := range values{
			keyToSend := &pb.Value{Value: value}
			if err := stream.Send(keyToSend); err != nil {
				return err
			}
		}
	} else {
		error := status.Error(codes.NotFound, fmt.Sprintf("Key %v not found", keySent))
		return error
	}
	return nil
}

func (kvs *keyValueServer) Put(ctx context.Context, message *pb.Message) (*pb.Empty, error){
	logOperation("Put")
	keySent := message.Key.GetKey()
	valueSent := message.Value.GetValue()
	memory[keySent] = []string{valueSent}
	return &pb.Empty{}, nil
}

func (kvs *keyValueServer) Append(ctx context.Context, message *pb.Message) (*pb.Empty, error){
	logOperation("Append")
	keySent := message.Key.GetKey()
	valueSent := message.Value.GetValue()
	memory[keySent] = append(memory[keySent], valueSent)
	return &pb.Empty{}, nil
}

func logOperation(operation string){
	log.Printf("Receiving %v operation\n", operation)
	log.Printf("Memory: %v\n", memory)
}

func newServer() *keyValueServer {
	server := &keyValueServer{}
	return server
}

func main(){
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Problem with the port: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDistributeStoreServer(grpcServer, newServer())
	log.Printf("Listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}