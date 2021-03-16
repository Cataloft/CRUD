package main

import (
	"log"
	"os"

	"awesomeProject/internal/handler"
	"awesomeProject/internal/server"
)

func main() {
	dsn := os.Getenv("DSN")

	h := handler.New(dsn)

	srv := server.New(h)
	log.Fatal(srv.Start())
}
