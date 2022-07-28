package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[string]string{}
	largest := 0
	// newString := ""
	if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s); i++ {

		if _, ok := m[string(s[i])]; !ok {
			m[string(s[i])] = string(s[i])
		} else {
			for k := range m {
				delete(m, k)
			}
			i--
		}
		if len(m) > largest {
			largest = len(m)
		}
	}
	fmt.Println(m, largest)

	return largest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
