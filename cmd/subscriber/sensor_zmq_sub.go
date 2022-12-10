package main

import (
	"fmt"
	"goFleet/internal/adapters/zmq"
	"goFleet/internal/application/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("application is starting...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	go exec()     // runs until main() exits
	<-quitChannel // block until quit is triggered

	log.Println("application is closing...")
}

func env(name string) (string, error) {
	variable, exists := os.LookupEnv(name)
	if !exists {
		return variable, fmt.Errorf("environment variable must be set: %s", name)
	}
	return variable, nil
}

func exec() {
	zmqSubEndpoint, err := env("ZMQ_SUB_ENDPOINT")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}
	topic, err := env("TOPIC")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}

	zmqSubAdapter, err := zmq.NewSubAdapter(zmqSubEndpoint, topic)
	if err != nil {
		log.Fatalf("failed to create adapter: %v", err)
	}
	defer zmqSubAdapter.Close()

	saveEndpoint, err := env("SAVE_ENDPOINT")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}
	subscriptionHandler, err := api.NewEndpointWriter(saveEndpoint)
	if err != nil {
		log.Fatalf("failed to create subscription handler: %v", err)
	}
	defer subscriptionHandler.Close()

	application := api.NewSubApplication(zmqSubAdapter, subscriptionHandler)

	for {
		err := application.ReceiveAndHandle()
		if err != nil {
			log.Fatalf("failed to receive or handle: %v", err)
		}
	}
}
