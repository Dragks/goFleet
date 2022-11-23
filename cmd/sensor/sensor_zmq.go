package main

import (
	"fmt"
	"goFleet/internal/adapters/zmq"
	"goFleet/internal/application/api"
	"goFleet/internal/application/core/sensor"
	"log"
	"os"
)

func main() {
	var err error
	fmt.Println("application is starting...")

	zmqConnection := os.Getenv("ZMQ_CONNECTION")
	sensorId := os.Getenv("SENSOR_ID")

	zmqAdapter, err := zmq.NewAdapter(zmqConnection)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer zmqAdapter.CloseConnection()

	sens := sensor.New(sensorId)
	applicationApi := api.NewZmqApplication(zmqAdapter, sens)

	val, err := applicationApi.GetRead()
	if err != nil {
		log.Fatalf("failed to get sensor read: %v", err)
	}
	fmt.Printf("got sensor read with value: %v\n", val)
	fmt.Println("application is closing...")
}
