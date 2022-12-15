package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	cache "memcache/api/proto"
	"os"
	"time"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conn, err := grpc.Dial(os.Getenv("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := cache.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(c.Set(ctx, &cache.Item{Key: "1", Value: "value"}))

	fmt.Println("-----")

	fmt.Println(c.Get(ctx, &cache.Key{Key: "1"}))
}
