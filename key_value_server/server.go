package main

import (
	"context"
	pb "distribute_store/rpc_definition"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	)

var (
	port = flag.Int("port", 8000, "Sever Port")
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
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDistributeStoreServer(grpcServer, newServer())
	log.Printf("Listening on localhost:%d\n", *port)
	grpcServer.Serve(lis)
}