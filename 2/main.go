package main

import "fmt"

func main() {
	fmt.Println(dissimilarity("bxcn", "abncx"))
	fmt.Println(dissimilarity("", "y"))
	fmt.Println(dissimilarity("annqalff", "fqlnannaf"))

}

func dissimilarity(s, t string) (output string) {
	mp := make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}
	mp2 := make(map[rune]int)
	for _, v := range t {
		mp2[v]++
	}
	for i, v := range mp2 {
		if mp[i] != v {
			output += string(i)
		}
	}
	return output
}
