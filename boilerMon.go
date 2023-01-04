package main

import (
	"fmt"
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

	err := boiler.AddZone("LivingRoom", 2)
	if err != nil {
		fmt.Println("Error adding LivingRoom zone")
	}
	err = boiler.AddZone("BedRoom", 3)
	if err != nil {
		fmt.Println("Error adding bedroom zone")
	}
	err = boiler.AddZone("Basement", 4)
	if err != nil {
		fmt.Println("Error adding Basement zone")
	}

	for {
		if err = boiler.Update(); err != nil {
			fmt.Printf("Error updating boiler: %v\n", err)
		} else {
			boiler.PrintJson()
		}
		fmt.Println("**************************************")
		time.Sleep(5 * time.Second)
	}

}
