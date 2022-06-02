package main

func check(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for _, r := range s {
		v := r - 'a'
		m[v] += 1
	}
	n := [26]int{}
	for _, r := range t {
		v := r - 'a'
		n[v] += 1
	}
	return m == n
}

func groupAnagrams(strs []string) [][]string {
	var matrix [][]string
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {

		}
	}
	return matrix
}
