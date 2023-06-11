package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumaniyozov/glog/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Blog service Started")

	//TODO: Move to environment variable on production
	//ports:
	//	- "5433:5432"
	//environment:
	//POSTGRES_USER: postgres
	//POSTGRES_PASSWORD: password
	//POSTGRES_DB: glog
	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres:password@localhost:5433/glog")
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatalf("something went wrong connecting to db: %v\n", err)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	proto.RegisterBlogServiceServer(s, &server{})

	go func() {
		log.Println("Starting grpc blog server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	log.Println("Stopping the server")
	s.Stop()
	log.Println("Closing the listener")
	lis.Close()
	log.Println("End of Program")
}
