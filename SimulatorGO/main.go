package main

import (
	"fmt"

	route2 "github.com/code/imersao-full-cycle-simulator/application/route"
)

func main() {
	route := route2.Route{
		Id:       "1",
		ClientId: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Print(stringJson[0])
}
