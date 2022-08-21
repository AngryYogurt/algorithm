package common

import "math/rand"

// RandomIntSlice 随机数列表
func RandomIntSlice(rangeMin, rangeMax, size int, cr *rand.Rand) []int {
	result := make([]int, size)
	ra := r
	if cr != nil {
		ra = cr
	}
	for i, _ := range result {
		result[i] = ra.Intn(rangeMax-rangeMin+1) + rangeMin
	}
	return result
}

// ShuffledContinuousIntSlice 打乱的连续自然数列表
func ShuffledContinuousIntSlice(size int) []int {
	result := make([]int, size)
	for i, _ := range result {
		result[i] = i
	}
	r.Shuffle(size, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
