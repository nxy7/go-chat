package main

import (
	"fmt"
	"go-chat/server"
)

func main() {
	s := server.Server{Port: 8282, Prefix: "/api"}
	err := s.Start()
	if err != nil {
		fmt.Println(err)
	}
}
