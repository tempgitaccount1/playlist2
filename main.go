package main

import (
	"fmt"
	"playlist/server"
)

func main() {
	fmt.Println("start server!!")
	s := server.New()
	s.Start()
	fmt.Println("did the thing")
}
