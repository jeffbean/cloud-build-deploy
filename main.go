package main

import (
	"fmt"

	"gopkg.in/redis.v3"
)

func main() {
	ExampleNewClient()
	fmt.Print("Hello")
}
func ExampleNewClient() {
	host := "localhost"
	port := "6379"
	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
