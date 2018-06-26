package main

import (
	"sync"

	"testing"
)

var (
	size       = 100000
	nbParallel = 16

	mux = &sync.Mutex{}

	storeDataChan         []int
	storeDataChanParallel []int
	storeDataMutex        []int
)

func init() {
	storeDataChan = make([]int, size)
	storeDataChanParallel = make([]int, size)
	storeDataMutex = make([]int, size)
}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := make(chan int)
		go func() {
			for i := 0; i < size; i++ {
				ss <- i
			}
			close(ss)
		}()

		for {
			val, ok := <-ss
			if !ok {
				break
			}

			storeDataChan[val] = val
		}
	}
}

func BenchmarkChanParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := make(chan int)
		go func() {
			for i := 0; i < size; i++ {
				ss <- i
			}
			close(ss)
		}()

		var wg sync.WaitGroup
		fct := func(w *sync.WaitGroup) {
			for {
				val, ok := <-ss
				if !ok {
					break
				}

				storeDataChanParallel[val] = val
			}
			w.Done()
		}

		wg.Add(nbParallel)
		for i := 0; i < nbParallel; i++ {
			go fct(&wg)
		}
		wg.Wait()
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			mux.Lock()
			storeDataMutex[i] = i
			mux.Unlock()
		}
	}
}
