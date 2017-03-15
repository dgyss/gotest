package basic

import (
	"sync"
	"testing"
)

func Defer()  {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
}

func NoDefer()  {
	var lock sync.Mutex
	lock.Lock()
	lock.Unlock()
}

func BenchmarkDefer(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Defer()
	}
}

func BenchmarkNoDefer(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		NoDefer()
	}
}

