package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/seaskythe/TChatP/client"
	"github.com/seaskythe/TChatP/server"
)

func main() {
	mode := flag.String("mode", "", "Mode to run: 'client' or 'server'")
	address := flag.String("address", "localhost:9092", "Address to connect to or listen on")
	flag.Parse()

	switch *mode {
	case "client":
		runClient(*address)
	case "server":
		runServer(*address)
	default:
		fmt.Println("Invalid mode. Use 'client' or 'server'.")
		flag.Usage()
		os.Exit(1)
	}
}

func runClient(address string) {
	fmt.Printf("Connecting to TChatP server at %s \n\n", address)
	client.ConnectToServer(address)
}

func runServer(address string) {
	fmt.Printf("Starting TChatP server at %s \n\n", address)
	go func() {
		server.ListenConn(address)
	}()

	// Handle graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Shutting down server...")
	// Call server shutdown logic if any
	// server.Shutdown() // Implement this function in your server package
	fmt.Println("Server gracefully stopped.")
}
