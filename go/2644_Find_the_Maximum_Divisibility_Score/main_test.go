package leetcode

import "testing"

func TestMaxDivScore(t *testing.T) {
	t.Run("Case: 1", func(t *testing.T) {
		nums := []int{31, 91, 47, 6, 37, 62, 72, 42, 85}
		divisors := []int{95, 10, 8, 43, 21, 63, 26, 45, 23, 69, 16, 99, 92, 5, 97, 69, 33, 44, 8}
		got := maxDivScore(nums, divisors)
		want := 5

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Case: 2", func(t *testing.T) {
		nums := []int{39, 70, 33}
		divisors := []int{27, 73, 37, 31, 28, 82, 96, 12, 31, 77, 17, 83, 19, 46, 7, 4, 74, 70, 66, 73, 25, 50, 79}
		got := maxDivScore(nums, divisors)
		want := 7

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
