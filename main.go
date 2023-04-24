package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()

	for i := -7; i <= 7; i++ {
		fmt.Println("date: ", date.AddDate(0,0,i))
	}
}