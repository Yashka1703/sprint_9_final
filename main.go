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

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, fmt.Errorf("wrong slice size: %v", size)
	}

	randMap := make([]int, size)
	rand.NewSource(time.Now().UnixNano())

	for i := range randMap {
		randMap[i] = rand.Int()
	}
	if len(randMap) != size {
		return nil, fmt.Errorf("wrong slice size: %v", size)
	}
	return randMap, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}

	maxInt := data[0]

	for _, p := range data[1:] {
		if maxInt < p {
			maxInt = p
		}
	}
	return maxInt, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	sizeChunk := len(data) / CHUNKS
	res := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * sizeChunk
		end := start + sizeChunk

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(start, end int, i int) {
			defer wg.Done()
			maxInChunk, _ := maximum(data[start:end])
			res[i] = maxInChunk
		}(start, end, i)
	}

	wg.Wait()

	maxInChunks, _ := maximum(res)
	return maxInChunks
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)
	nums, _ := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, _ := maximum(nums)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(nums)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
