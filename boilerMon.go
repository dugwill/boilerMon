package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dugwill/triangleTube"
)

var boiler *triangleTube.TriangleTube
var err error

func main() {

	if boiler, err = triangleTube.NewBoiler(1, "/dev/ttyUSB0"); err != nil {
		fmt.Printf("Error creating new boiler instance: %v", err)
		return
	}

	err := boiler.AddZone("LivingRoom", 4)
	if err != nil {
		fmt.Println("Error adding LivingRoom zone")
	}
	err = boiler.AddZone("BedRoom", 23)
	if err != nil {
		fmt.Println("Error adding bedroom zone")
	}
	err = boiler.AddZone("Basement", 25)
	if err != nil {
		fmt.Println("Error adding Basement zone")
	}

	// Open connection to ELK via Vector
	eventEndpoint := "127.0.0.1:9000"
	eventSocket, err := net.Dial("tcp", eventEndpoint)
	if err != nil {
		log.Println("Error opening event socket. Events will not be stored")
	} else {
		log.Println("StoreEvents enabled")
	}

	for {
		if err = boiler.Update(); err != nil {
			fmt.Printf("Error updating boiler: %v\n", err)
		} else {
			boiler.PrintJson()

			jsonString, err := json.Marshal(boiler)
			if err != nil {
				log.Printf("error marshalling event %v ", err)
				continue
			}
			fmt.Fprintln(eventSocket, bytes.NewBuffer(jsonString))
		}
		fmt.Println("**************************************")
		time.Sleep(5 * time.Second)
	}

}
