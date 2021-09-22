package main

import (
	"github.com/maanxester/webapi-golang/database"
	"github.com/maanxester/webapi-golang/server"
)

func main() {
	database.RunDB()

	server := server.NewServer()

	server.Run()
}
