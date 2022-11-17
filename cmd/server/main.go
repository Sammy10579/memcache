package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	cache "memcache/api/proto"
	"memcache/pkg/service"
	"net"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("Np .env file found")
	}

	s := grpc.NewServer()
	srv := &service.CacheServer{}
	cache.RegisterCacheServer(s, srv)

	l, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting gRPC listener on port " + os.Getenv("PORT"))
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
