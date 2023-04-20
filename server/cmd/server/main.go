package main

import (
	"log"
	"net"

	"github.com/Lalipopp4/grpc-uni/server/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lst, err := net.Listen("tcp", ":81")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running...")
	s := grpc.NewServer()
	reflection.Register(s)
	handler.InitServer(s)
	handler.Maker()
	if err := s.Serve(lst); err != nil {
		log.Fatalln(err)
	}
}
