package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sf := strings.Fields(s)

	wc := make(map[string]int)

	for _, w := range sf {
		wc[w] += 1
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}
