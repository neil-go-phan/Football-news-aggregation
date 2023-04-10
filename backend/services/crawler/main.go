package main

import (
	"crawler/handlers"
	"crawler/helper"
	"crawler/services"
	"log"
)

func main() {
	env, err := crawlerhelpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	services.CrawlMatchDetail("/truc-tiep-ket-qua/west-ham-vs-newcastle-113358.html")
	handlers.GRPCServer(env.GRPCPort)
} 

