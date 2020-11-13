package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func normalize(word string) (neWord string) {
	wordNew := strings.ToLower(word)
	m := regexp.MustCompile("[,|.|!|\\n]+?")
	wordNew = m.ReplaceAllString(word, "")

	return wordNew
}

func findWord(sentence string) (word map[string]int) {
	newArray := strings.Split(sentence, " ")
	dicWords := make(map[string]int)

	for _, word := range newArray {
		norWord := normalize(word)

		_, ok := dicWords[norWord]
		if ok {
			dicWords[norWord]++
		} else {
			dicWords[norWord] = 1
		}
	}
	return dicWords
}

func main() {
	fmt.Printf("Write a sentence and I'll find the most repeated word: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	word := findWord(sentence)
	fmt.Println(word)
}
