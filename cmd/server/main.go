package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"

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

	connectionChannel, err := connection.Channel()

	if err != nil {
		fmt.Printf("Error getting connection channel of type AMPQP Channel")
	}

	pubsub.PublishJSON(connectionChannel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel
	fmt.Println("Shutting down")

}
