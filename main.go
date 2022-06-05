package main

import (
	"math"
)

func containsDuplicate(nums []int) bool {
	// Map of numbers
	m := make(map[int]int)

	for _, i := range nums {
		// If the number is a duplicate return early
		if _, exists := m[i]; exists {
			return true
		}

		m[i] = i
	}
	return false
}

func isAnagram(s, t string) bool {
	// Map of how many times a character has been seen
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	// Populate map
	for _, char := range s {
		map1[char] += 1
	}
	for _, char := range t {
		map2[char] += 1
	}

	// Compare values in both maps
	for char, count := range map1 {
		if map2[char] != count {
			return false
		}
	}

	return true
}

func twoSum(nums []int, target int) []int {
	// Map of numbers to list of indexes
	var indexOf = make(map[int][]int, len(nums))

	// Populate Map
	for i, n := range nums {
		indexOf[n] = append(indexOf[n], i)
	}

	for index1, n := range nums {
		// If wanted number exists in map
		if val, ok := indexOf[target-n]; ok {
			for _, index2 := range val {
				// if index is not a duplicate
				if index2 != index1 {
					// Return early
					return []int{index1, index2}
				}
			}
		}
	}

	return []int{0, 0}
}

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0, k)
	freq := make(map[int]int, len(nums))
	m2 := make(map[int][]int, len(nums))

	for _, num := range nums {
		freq[num] += 1
	}

	for k, v := range freq {
		m2[v] = append(m2[v], k)
	}

	x := 0
	for i := len(nums); i > 0; i-- {
		if x == k {
			return ans
		}

		if val, ok := m2[i]; ok {
			for _, v := range val {
				ans = append(ans, v)
				x++
			}
		}
	}

	return ans
}

// Solution with for loops & slices
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	count := len(nums) - 1

	ans[0] = 1
	ans[count] = 1

	// Add prefix to answer array
	for i := 1; i <= count; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	// Temporary variable to store postfix
	postfix := 1
	for i := count - 1; i >= 0; i-- {
		// Calculate postfix
		postfix = nums[i+1] * postfix
		// Multiply Prefix & Postfix
		ans[i] *= postfix
	}
	return ans
}

// Solution using hashtables
func productExceptSelf_v1(nums []int) []int {
	freq := make(map[int]int)
	ans := make([]int, len(nums))

	for _, n := range nums {
		freq[n] += 1
	}

	for i, n := range nums {
		ans[i] = 1
		for k, v := range freq {
			if k != n {
				ans[i] *= powInt(k, v)
			} else if (k == n) && (v > 1) {
				ans[i] *= powInt(k, v-1)
			}
		}
	}
	return ans
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
