package main

import (
	"fmt"
	"goFleet/internal/application/api"
	"goFleet/internal/application/core/sensor"
	"log"
)

func main() {
	var err error
	fmt.Println("application is starting...")

	sen := sensor.New()
	applicationApi := api.NewApplication(sen)

	val, err := applicationApi.GetRead()
	if err != nil {
		log.Fatalf("failed to get sensor read: %v", err)
	}
	fmt.Printf("got sensor read with value: %v\n", val)
	fmt.Println("application is closing...")
}
