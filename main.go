package main

import (
	"github.com/lekoller/api-with-go/database"
	"github.com/lekoller/api-with-go/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
