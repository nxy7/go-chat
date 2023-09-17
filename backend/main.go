package main

import (
	"fmt"
	"go-chat/config"
	"go-chat/server"
	"log"
)

func main() {
	config := config.FromEnv()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Config loaded")

	if err := server.Start(config); err != nil {
		fmt.Println(err)
	}
}
