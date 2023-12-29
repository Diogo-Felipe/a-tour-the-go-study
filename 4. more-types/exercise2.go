package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	stringSlice := strings.Fields(s)

	var wordsMap map[string]int
	wordsMap = make(map[string]int)
	for _, value := range stringSlice {
		_, exist := wordsMap[value]

		if exist {
			wordsMap[value] = wordsMap[value] + 1
		} else {
			wordsMap[value] = 1
		}
	}

	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
