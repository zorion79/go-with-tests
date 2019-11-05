package concurrency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func MockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func SlowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsite(MockWebsiteChecker, websites)

	assert.Equal(t, want, got)
}

func BenchmarkSlowCheckWebsite(b *testing.B) {
	urls := prepareUrls()

	for i := 0; i < b.N; i++ {
		CheckWebsite(SlowStubWebsiteChecker, urls)
	}
}

func BenchmarkFastCheckWebsite(b *testing.B) {
	urls := prepareUrls()

	for i := 0; i < b.N; i++ {
		CheckWebsite(MockWebsiteChecker, urls)
	}
}

func prepareUrls() []string {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	return urls
}
