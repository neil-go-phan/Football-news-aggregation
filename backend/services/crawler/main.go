package main

import (
	"backend/services/crawler/grpcServer"
)


func main() {
	// articles, _ := crawl.SearchKeyWord("Ngoại hạng anh")
	// for index, a := range articles {
	// 	fmt.Println("index: ", index, " title: ", a.Title)
	// }
	grpcserver.GRPCServer()
}
