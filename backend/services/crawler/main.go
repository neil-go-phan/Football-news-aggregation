package main

import (
	"crawler/handlers"
	"crawler/helper"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)
func init() {
	// setup log
	log.SetLevel(log.InfoLevel)
	format := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	}
	log.SetFormatter(format)

}

func main() {
	env, err := crawlerhelpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}

	logFile, err := os.OpenFile("crawlerlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "crawlerlog.log")
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	handlers.GRPCServer(env.GRPCPort)
} 
