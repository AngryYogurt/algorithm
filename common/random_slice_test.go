package common

import (
	"fmt"
	"testing"
)

func TestRandomIntSlice(t *testing.T) {
	var result []int
	result = RandomIntSlice(47, 80, 60)
	fmt.Println(result)
}

func TestShuffledContinuousIntSlice(t *testing.T) {
	var result []int
	result = ShuffledContinuousIntSlice(101)
	var sum int
	for _, v := range result {
		sum += v
	}

	fmt.Println(sum)
	fmt.Println(result)
}
