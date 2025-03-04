package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	frase := strings.Fields(s)
	m := map[string]int{}

	for _, valor := range frase {
		m[valor] += 1
	}
	return m
}

// func main() {
// 	wc.Test(WordCount)
// }
