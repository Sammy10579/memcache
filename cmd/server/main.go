package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	cache "memcache/api/proto"
	"memcache/pkg/service"
	"memcache/pkg/storage"
	"net"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	mc := memcache.New("localhost:11211")
	st := storage.NewStorage(mc)
	csrv := service.NewCacheServer(st)
	s := grpc.NewServer()
	cache.RegisterCacheServer(s, csrv)

	l, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting gRPC listener on port " + os.Getenv("PORT"))
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
