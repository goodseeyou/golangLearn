package main

import (
	"net/http"
	"sync"
	"testing"
)

// net http client is thread safe
// this is better for performance
func BenchmarkOneClientRequest(b *testing.B) {
	client := &http.Client{}
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client.Get("http://localhost:8888/")
		}()
	}

	wg.Wait()
}

func BenchmarkMultipleClientRequest(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}
			client.Get("http://localhost:8888/")
		}()
	}

	wg.Wait()
}
