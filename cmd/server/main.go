package main

import (
	"fmt"
	"log"
)

func Run() error {
	fmt.Println("Woking...")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal("Something went wrong.")
	}
}
