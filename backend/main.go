package main

import (
	"fmt"
	"go-chat/server"
)

func main() {
	s := server.Server{Port: 8282}
	err := s.Start()
	if err != nil {
		fmt.Println(err)
	}
}
