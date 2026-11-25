package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns size random positive integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for index := range data {
		data[index] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

// maximum returns the largest value, or zero if data is empty.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

// maxChunks finds the largest value using up to CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	workers := min(CHUNKS, len(data))
	chunkSize := len(data) / workers
	remainder := len(data) % workers
	maxima := make([]int, workers)
	var group sync.WaitGroup
	group.Add(workers)

	start := 0
	for index := range maxima {
		end := start + chunkSize
		if index < remainder {
			end++
		}
		go func(index int, chunk []int) {
			defer group.Done()
			maxima[index] = maximum(chunk)
		}(index, data[start:end])
		start = end
	}

	group.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	started := time.Now()
	sequentialMax := maximum(data)
	sequentialTime := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", sequentialMax, sequentialTime)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	started = time.Now()
	parallelMax := maxChunks(data)
	parallelTime := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", parallelMax, parallelTime)
}
