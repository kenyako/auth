package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "github.com/kenyako/auth/pkg/auth_v1"
)

const (
	address = "localhost:50051"
	userID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewUserAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf("User info:\nID: %d\nName: %s\nEmail: %s\nRole: %s\nCreated at: %v\n Updated at: %v",
		r.GetId(), r.GetName(), r.GetEmail(), r.GetRole(),
		r.GetCreatedAt(), r.GetUpdatedAt())
}
