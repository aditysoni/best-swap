package main

// Import necessary packages for the main application
import (
	"log"

	"github.com/joho/godotenv"
	"github.com/aditysoni/flashloan-bot/server"

)

func main() {
	// Connect to Ethereum node
	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file", err)
    }
 
	server.RunServer()

}
