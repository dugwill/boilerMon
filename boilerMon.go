package main

import (
	"fmt"

	"github.com/dugwill/triangleTube"
)

var boiler *triangleTube.TriangleTube
var err error

func main() {

	if boiler, err = triangleTube.NewBoiler(1, "com3"); err != nil {
		fmt.Printf("Error creating new boiler instance: %v", err)
		return
	}

	//fmt.Println("Slave ID: ", boiler.ReportSlaveID())

	if err = boiler.Update(); err != nil {
		fmt.Printf("Error updating boiler: %v\n", err)
	} else {
		fmt.Printf("Boiler Temp: %v\n", boiler.BoilerSupplyTemp)
		fmt.Printf("Boiler Return Temp: %v\n", boiler.BoilerReturnTemp)

	}

}
