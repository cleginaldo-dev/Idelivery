package main

import (
	"log"

	"github.com/code/imersao-full-cycle-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("Hola", "readtest", producer)

	for {
		_ = 1
	}

	// route := route2.Route{
	// 	Id:       "1",
	// 	ClientId: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Print(stringJson[0])
}
