package durstenfeld

import (
	"testing"
	"sort"
	"durstenfeld"
)

func TestSize(t *testing.T) {
	nums := durstenfeld.RandInts(0, 10)
	if len(nums) != 10 {
		t.Error("Expected 10 numbers for range 0..10", nums)
	}
}

func TestLow(t *testing.T) {
	nums := durstenfeld.RandInts(10, 20)
	lo := nums[0]
	for _, n := range nums {
		if n < lo {
			lo = n
		}
	}
	if lo != 10 {
		t.Error("Expected lowest number to be 10", nums)
	}
}

func TestHigh(t *testing.T) {
	nums := durstenfeld.RandInts(50, 70)
	hi := nums[0]
	for _, n := range nums {
		if n > hi {
			hi = n
		}
	}
	if hi != 69 {
		t.Error("Expected lowest number to be 69", nums)
	}
}

func TestUnique(t *testing.T) {
	nums := durstenfeld.RandInts(0, 10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(nums)
	if !func() bool {
		for i, n := range nums {
			if n != expected[i] {
				return false
			}
		}
		return true
	}()	{
		t.Error("Expected numbers to be unique", nums)
	}
}

func TestShuffleness(t *testing.T) {
	unshuffled := make([]int, 30)
	for i := 0; i < 30; i++ {
		nums := durstenfeld.RandInts(0, 30)
		for j, n := range nums {
			if j == n {
				unshuffled[j]++
			}
		}
	}
	sum := 0
	for _, n := range unshuffled {
		sum += n
	}
	if sum > 40 {
		t.Error("High unshuffled count", sum)
	}
}
