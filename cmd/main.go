package main

import (
	"fmt"
	"goFleet/internal/adapters"
	"goFleet/internal/application/api"
	"goFleet/internal/application/core/sensor"
	"log"
	"os"
)

func main() {
	var err error
	fmt.Println("application is starting...")

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	sensorId := os.Getenv("SENSOR_ID")

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseConnection()

	core := sensor.New(sensorId)
	applicationApi := api.NewApplication(dbAdapter, core)

	val, err := applicationApi.GetRead()
	if err != nil {
		log.Fatalf("failed to get sensor read: %v", err)
	}
	fmt.Printf("got sensor read with value: %v\n", val)
	fmt.Println("application is closing...")
}
