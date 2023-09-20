package main

func signature(s string) [26]int {
	var sig [26]int
	for _, i := range s {
		sig[i-'a']++
	}
	return sig
}

func groupAnagrams(strs []string) [][]string {
	signatures := make(map[[26]int][]string)
	for _, s := range strs {
		sig := signature(s)
		signatures[sig] = append(signatures[sig], s)
	}
	var groups [][]string
	for _, g := range signatures {
		groups = append(groups, g)
	}
	return groups
}
