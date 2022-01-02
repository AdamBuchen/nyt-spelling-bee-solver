package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const wordsFile = "./words.txt"
const minLength = 4

var words []string = parseWordList()
var inner []string
var outer string

func main() {

	innerPtr := flag.String("inner", "", "The single letter that must be included in every word")
	outerPtr := flag.String("outer", "", "Single string containing all the outer characters")

	flag.Parse()

	//inner
	inner := *innerPtr
	if len(inner) != 1 {
		panic("Inner must be a single character")
	}
	inner = strings.ToLower(inner)

	//outer := make([]string, 0)
	outerStr := *outerPtr
	outer := make([]string, 0)
	chars := []rune(outerStr)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		char = strings.ToLower(char)
		outer = append(outer, char)
	}

	if len(outer) < 1 || len(outer) > 10 {
		panic("Outer must contain between 1 and 10 characters")
	}

	matchingWords := getMatchingWords(inner, outer)
	for _, matchedWord := range matchingWords {
		fmt.Println(matchedWord)
	}

}

func getMatchingWords(inner string, outer []string) []string {

	matchedWords := make(map[int][]string)
	sorted := make([]string, 0)

	var innerMatched bool
	var outerMatched bool
	var allMatched bool
	var wordLength int
	maxLength := 0

	outerLen := len(outer)

	for _, word := range words { //Loop through all dictionary words
		innerMatched = false
		allMatched = true
		wordLength = len(word)
		if wordLength < minLength {
			continue
		}

		for i := 0; i < wordLength; i++ { //Looping through all letters of the word
			char := string(word[i])
			char = strings.ToLower(char)
			if inner == char { //char matches my inner
				innerMatched = true
				continue
			}
			outerMatched = false
			for j := 0; j < outerLen; j++ { //Looping through all letters of outer
				if outer[j] == char { //a letter from the word matches my outer
					outerMatched = true
					break
				}
			}
			if !outerMatched {
				allMatched = false
				break
			}
		}
		//Done looping over letters for the word
		if allMatched && innerMatched {
			if _, ok := matchedWords[wordLength]; !ok {
				matchedWords[wordLength] = make([]string, 0)
			}
			matchedWords[wordLength] = append(matchedWords[wordLength], word)
			if wordLength > maxLength {
				maxLength = wordLength
			}
		}
	}

	for k := maxLength; k > 0; k-- {
		if _, ok := matchedWords[k]; ok {
			for _, word := range matchedWords[k] {
				sorted = append(sorted, word)
			}
		}
	}

	return sorted
}

func parseWordList() []string {
	file, err := os.Open(wordsFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
