package main

import (
	"crawler/crawl"
	"fmt"
)


func main() {
	articles, _ := crawl.SearchKeyWord("Ngoại hạng anh")
	for index, a := range articles {
		fmt.Println("index: ", index, " title: ", a.Title)
	}

}
