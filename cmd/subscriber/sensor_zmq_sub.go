package main

import (
	"fmt"
	"goFleet/internal/adapters/db"
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
	var err error

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

	dbaseDriver, err := env("DB_DRIVER")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}
	dsourceName, err := env("DS_NAME")
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.Close()

	application := api.NewSubApplication(zmqSubAdapter, dbAdapter)

	for {
		val, err := application.SubscribeAndSave()
		if err != nil {
			log.Fatalf("failed to get sensor read: %v", err)
		}
		log.Printf("got sensor read with value: %v\n", val)
	}
}
