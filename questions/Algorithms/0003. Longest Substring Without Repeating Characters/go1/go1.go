package main

import (
	"fmt"
	"math"
)

// ------------------------------- solution begin -------------------------------
func lengthOfLongestSubstring(s string) int {
	max := 0
	n := len(s)
	start := 0
	end := 0
	mp := make(map[string]int, 0)
	if n < 2 {
		return n
	}
	for ; end < n; end++ {
		char := s[end : end+1]
		v, ok := mp[char]
		if ok {
			start = int(math.Max(float64(start), float64(v)))
		}
		max = int(math.Max(float64(max), float64(end-start+1)))
		mp[char] = end + 1
	}
	return max
}

// ------------------------------- solution end ---------------------------------

func main() {
	var str = "abcabcbb"
	fmt.Printf("Input:  %s\n", str)
	fmt.Printf("Output: %d\n", lengthOfLongestSubstring(str))
}
