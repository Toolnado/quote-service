package main

import (
	"context"
	"log"
	"net"

	"github.com/Toolnado/quote-service.git/api"
	"github.com/Toolnado/quote-service.git/internal/onederx"
	"github.com/Toolnado/quote-service.git/internal/rpc"
	"google.golang.org/grpc"
)

func main() {
	source := onederx.NewSource()
	source.Start(context.Background())

	service := rpc.NewService()
	service.AddSource(source)
	server := grpc.NewServer()
	api.RegisterQuotesServer(server, service)
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start listening...")

	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
