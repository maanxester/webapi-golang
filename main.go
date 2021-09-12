package main

import "github.com/maanxester/webapi-golang/server"

func main() {
	server := server.NewServer()

	server.Run()
}
