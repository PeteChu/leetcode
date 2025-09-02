package leetcode

import (
	"sort"
)

func maxDivScore(nums []int, divisors []int) int {
	sort.Ints(divisors)
	x := divisors[0]
	m := make(map[int]int)
	for _, i := range divisors {
		m[i] = 0
		for _, j := range nums {
			if j%i == 0 {
				m[i] += 1
			}
		}
	}

	score := m[x]
	for i, v := range m {
		if v > score {
			score = v
			x = i
		}
		if v == score && i < x {
			x = i
		}
	}
	return x
}
