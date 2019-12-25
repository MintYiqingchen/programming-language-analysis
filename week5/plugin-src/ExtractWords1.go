package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func loadStopWords() map[string]bool {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	words := strings.Split(strings.TrimSpace(string(content)), ",")
	res := map[string]bool{}
	for _, w := range words {
		res[w] = true
	}
	return res
}
func ExtractWords(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	regp, _ := regexp.Compile(`[\W_]+`)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	wordlist := strings.Fields(strings.ToLower(string(content)))
	res := []string{}
	stopwords := loadStopWords()
	for _, w := range wordlist {
		if _, ok := stopwords[w]; !ok && len(w) > 1 {
			res = append(res, w)
		}
	}
	return res
}
