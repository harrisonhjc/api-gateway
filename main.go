package main

import (
	"api-gateway/gateway"
	"log"
)

func main() {
	log.Println("running ...........")
	starter := gateway.NewGatewayStarter()
	log.Println("running Start...........")
	starter.Start()
}
