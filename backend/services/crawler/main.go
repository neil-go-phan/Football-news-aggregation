package main

import (
	"crawler/handlers"
	"crawler/helper"
	"log"
)

func main() {
	env, err := crawlerhelpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	handlers.GRPCServer(env.GRPCPort)
} 

