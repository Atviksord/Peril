package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	connectionString := "amqp://guest:guest@localhost:5672"
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		fmt.Println("Couldnt establish connection", err)
	}

	defer connection.Close()
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel
	fmt.Println("Shutting down")

}
