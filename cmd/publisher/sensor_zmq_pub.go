package main

import (
	"fmt"
	"goFleet/internal/adapters/zmq"
	"goFleet/internal/application/api"
	"goFleet/internal/application/core/sensor"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("application is starting...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	go exec()     // runs until main() exits
	<-quitChannel // block until quit is triggered

	log.Println("application is closing...")
}

func exec() {
	zmqPubEndpoint, err := env("ZMQ_PUB_ENDPOINT")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}
	topic, err := env("TOPIC")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}

	zmqPubAdapter, err := zmq.NewPubAdapter(zmqPubEndpoint)
	if err != nil {
		log.Fatalf("failed to create adapter: %v", err)
	}
	defer zmqPubAdapter.Close()

	sens, _ := sensor.NewMock()
	application := api.NewPubApplication(zmqPubAdapter, sens, topic)

	for {
		err := application.ReadAndPublish()
		if err != nil {
			log.Printf("failed to get sensor read: %v", err)
		}
		time.Sleep(time.Second)
	}
}

func env(name string) (string, error) {
	variable, exists := os.LookupEnv(name)
	if !exists {
		return variable, fmt.Errorf("environment variable must be set: %s", name)
	}
	return variable, nil
}
