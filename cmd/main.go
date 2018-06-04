package main

import (
	"fmt"
	"log"
	// "os"

	"github.com/Gin/interate-with-vue/server"
)

func main() {
	fmt.Println("hello world!")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	server := server.NewServer()
	server.Run()
}
