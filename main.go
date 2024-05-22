package main

import "os"

func main() {
	port := os.Getenv("PORT")

	server := newServer(port)
	server.Run()
}