package main

import (
	"log"

	"github.com/Lalipopp4/grpc-uni/client/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":81", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := handler.NewClient(conn)
	client.Request()
}
