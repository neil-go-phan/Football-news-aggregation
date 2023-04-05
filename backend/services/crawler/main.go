package main

import (
	"crawler/handlers"
	"crawler/helpers"
	"log"
)

func main() {
	env, err := helpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	handlers.GRPCServer(env.GRPCPort)
}

