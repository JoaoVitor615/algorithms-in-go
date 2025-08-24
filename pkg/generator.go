package pkg

import (
	"math/rand"
	"time"
)

// RandomGenerator provides utilities for generating random data
type RandomGenerator struct {
	rng *rand.Rand
}

// NewRandomGenerator creates a new RandomGenerator with current time seed
func NewRandomGenerator() *RandomGenerator {
	return &RandomGenerator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewRandomGeneratorWithSeed creates a new RandomGenerator with specific seed
func NewRandomGeneratorWithSeed(seed int64) *RandomGenerator {
	return &RandomGenerator{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// GenerateIntSlice generates a slice of random integers
func (rg *RandomGenerator) GenerateIntSlice(count, min, max int) []int {
	if count <= 0 {
		return []int{}
	}
	
	if min > max {
		min, max = max, min
	}
	
	result := make([]int, count)
	rangeSize := max - min + 1
	
	for i := 0; i < count; i++ {
		result[i] = rg.rng.Intn(rangeSize) + min
	}
	
	return result
}

// GenerateIntSliceDefault generates a slice with default range 1-1000
func (rg *RandomGenerator) GenerateIntSliceDefault(count int) []int {
	return rg.GenerateIntSlice(count, 1, 1000)
}

// GenerateSortedSlice generates a sorted slice of integers
func (rg *RandomGenerator) GenerateSortedSlice(count, min, max int) []int {
	slice := rg.GenerateIntSlice(count, min, max)
	// Simple insertion sort for small arrays
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
	return slice
}

// GenerateReverseSortedSlice generates a reverse sorted slice
func (rg *RandomGenerator) GenerateReverseSortedSlice(count, min, max int) []int {
	slice := rg.GenerateSortedSlice(count, min, max)
	// Reverse the slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// ShuffleSlice shuffles an existing slice in place
func (rg *RandomGenerator) ShuffleSlice(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rg.rng.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
