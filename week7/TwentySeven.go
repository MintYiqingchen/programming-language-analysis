package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

func characters(filename string) chan rune {
	res := make(chan rune)
	go func() {
		content, _ := ioutil.ReadFile(filename)
		for _, r := range bytes.Runes(content) {
			res <- r
		}
		close(res)
	}()
	return res
}

func allWords(filename string) chan string {
	charGen := characters(filename)
	res := make(chan string)
	go func() {
		startChar := true
		var word []rune
		for r := range charGen {
			if startChar {
				word = []rune{}
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					word = append(word, r)
					startChar = false
				}
			} else {
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					word = append(word, r)
				} else {
					startChar = true
					res <- strings.ToLower(string(word))
				}
			}
		}
		close(res)
	}()
	return res
}

func nonStopWords(filename string) chan string {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	text := strings.Split(strings.TrimSpace(string(content)), ",")
	stopwords := map[string]bool{}
	for i := 0; i < len(text); i++ {
		stopwords[text[i]] = true
	}
	wordsGen := allWords(filename)
	res := make(chan string)
	go func() {
		for w := range wordsGen {
			if _, ok := stopwords[w]; !ok && len(w) > 1 {
				res <- w
			}
		}
		close(res)
	}()
	return res
}

type Pair struct {
	word  string
	count int
}

func countAndSort(filename string) chan []Pair {
	wordsGen := nonStopWords(filename)
	res := make(chan []Pair)
	toPairSlice := func(freqs map[string]int) []Pair {
		ret := []Pair{}
		for k, v := range freqs {
			ret = append(ret, Pair{k, v})
		}
		sort.Slice(ret, func(i, j int) bool { return ret[i].count > ret[j].count })
		return ret
	}
	go func() {
		i := 1
		freqs := map[string]int{}
		for w := range wordsGen {
			if _, ok := freqs[w]; ok {
				freqs[w]++
			} else {
				freqs[w] = 1
			}
			if i%5000 == 0 {
				res <- toPairSlice(freqs)[:25]
			}
			i++
		}
		res <- toPairSlice(freqs)[:25]
		close(res)
	}()
	return res
}

func main() {
	for freqs := range countAndSort(os.Args[1]) {
		fmt.Println("-----------------------------")
		for i := range freqs {
			fmt.Println(freqs[i].word, " - ", freqs[i].count)
		}
	}
}
