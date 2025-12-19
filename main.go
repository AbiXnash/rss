package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Hello there!")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not defined.")
	}

	fmt.Println("Application is running on", port)
}
