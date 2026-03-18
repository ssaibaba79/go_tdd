package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCheckWebSites(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"http://bad.site.x",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"http://bad.site.x":          false,
	}

	mockChecker := func(url string) bool {
		return url != "http://bad.site.x"
	}

	got := CheckWebSites(mockChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%v, got %v", want, got)
	}
}

func BenchmarkNonBlockingCheckWebSites(b *testing.B) {
	urls := make([]string, 50)
	mockWebsiteChecker := func(url string) bool {
		time.Sleep(1 * time.Millisecond)
		return true
	}
	for i := range urls {
		urls[i] = fmt.Sprintf("http://test%dsite.com", i)
	}

	for b.Loop() {
		CheckWebSites(mockWebsiteChecker, urls)
	}
}
