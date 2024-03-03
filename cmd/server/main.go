package main

import (
	"log"

	"github.com/roblesdotdev/proglog/internal/server"
)

func Run() error {
	srv := server.NewHTTPServer(":8080")
	return srv.ListenAndServe()
}

func main() {
	if err := Run(); err != nil {
		log.Fatal("Something went wrong.")
	}
}
