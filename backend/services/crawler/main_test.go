package main

import "testing"

func BenchmarkSearchKeyWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchKeyWord("ngoai hang anh")
	}

}
// dùng goquery và chromedp đều tốn 5s để crawl 5 page
// goquery có thể crawl 100 page cùng lúc. vẫn với 5s
// chromedp không thể crawl quá 5 page cùng lúc
