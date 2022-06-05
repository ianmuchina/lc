package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwoSum(t *testing.T) {
	data := []struct {
		want  int
		input []int
		ans   []int
	}{
		{9, []int{2, 7, 11, 15}, []int{}},
		{6, []int{3, 2, 4}, []int{}},
		{6, []int{3, 3}, []int{}},
	}

	for _, d := range data {
		d.ans = twoSum(d.input, d.want)
		if d.want != (d.input[d.ans[1]] + d.input[d.ans[0]]) {
			t.Errorf("two Sum: %v  is not valid for %v", d.ans, d.want)
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		want  bool
		input []int
	}{
		{false, []int{0, 1, 2, 3, 4}},
		{false, []int{1, 2, 3, 4}},
		{true, []int{1, 2, 3, 1}},
		{true, []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	}

	for _, v := range tests {
		if containsDuplicate(v.input) != v.want {
			t.Error("Error in containsDuplicate")
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, v := range tests {
		if isAnagram(v.s, v.t) != v.want {
			t.Error("Error in containsDuplicate")
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	want := []int{24, 12, 8, 6}

	ans := productExceptSelf([]int{1, 2, 3, 4})
	if !cmp.Equal(ans, want) {
		t.Error("Error in Product Except Self")
	}

	ans = productExceptSelf_v1([]int{1, 2, 3, 4})
	if !cmp.Equal(ans, want) {
		t.Error("Error in Product Except Self v1")
	}
}
