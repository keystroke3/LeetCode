package main

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)
	charCount := 0
	for _, c := range s {
		if c == ' ' {
			continue
		}
		chars[c]++
		charCount++
	}
	for _, i := range t {
		if i == ' ' {
			continue
		}
		v, set := chars[i]
		if !set {
			return false
		}
		if v == 0 {
			return false
		}
		chars[i]--
		charCount--
	}
	return charCount == 0
}
