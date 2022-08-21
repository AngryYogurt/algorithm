package quick_sort

import (
	"fmt"
	"github.com/AngryYogurt/algorithm/common"
	"testing"
)

func orderCheck(src []int) bool {
	for idx, _ := range src {
		if idx >= 1 && src[idx] < src[idx-1] {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	var src []int
	for i := 0; i < 100; i++ {
		src = common.RandomIntSlice(0, 30, 100, nil)
		//src = common.ShuffledContinuousIntSlice(100)
		//src = []int{0, 5, 2, 4, 5, 3}
		fmt.Println("disorder=", src)
		QuickSort2(src)
		fmt.Println("ordered", src)
		if !orderCheck(src) {
			t.FailNow()
		}
	}
}

func BenchmarkQuickSort(t *testing.B) {
	var src []int
	for i := 0; i < 100; i++ {
		src = common.RandomIntSlice(0, 110, 100000, nil)
		QuickSort1(src)
	}
}

func BenchmarkQuickSortV2(t *testing.B) {
	var src []int
	for i := 0; i < 100; i++ {
		src = common.RandomIntSlice(0, 110, 100000, nil)
		QuickSort2(src)
	}
}

func BenchmarkQuickSortV3(t *testing.B) {
	var src []int
	for i := 0; i < 100; i++ {
		src = common.RandomIntSlice(0, 110, 100000, nil)
		QuickSort3(src)
	}
}
