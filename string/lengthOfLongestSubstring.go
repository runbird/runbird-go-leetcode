package string

import "math"

//leetcode 14 前缀和
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	dict := make(map[int32]int)
	j, max := 0, 0
	for i, v := range s {
		if _, ok := dict[v]; ok {
			j = int(math.Max(float64(j), float64(dict[v])+1))
		}
		dict[v] = i
		max = int(math.Max(float64(max), float64(i-j+1)))
	}
	return max
}

func main() {
	s := "tmmzuxt"
	print(lengthOfLongestSubstring(s))
}
