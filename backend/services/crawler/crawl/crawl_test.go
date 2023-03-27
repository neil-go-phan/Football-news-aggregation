package crawl

import "testing"

func BenchmarkSearchKeyWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchKeyWord("ngoai hang anh")
	}
}
