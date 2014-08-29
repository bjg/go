package durstenfeld

import (
	"math/rand"
	"time"
)

// Generate the specified range of randomly shuffled integers
func RandInts(start, finish int) []int {
	cnt := finish - start
	nums := make([]int, cnt)
	for i, _ := range nums {
		nums[i] = start + i
	}
	rand.Seed(time.Now().UnixNano())
	for i, _ := range nums {
		j := rand.Intn(len(nums))
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
