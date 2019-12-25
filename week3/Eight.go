package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type NextFunc func(interface{}, NextFunc) interface{}

func readFile(name interface{}, nextfunc NextFunc) interface{} {
	// name : string
	// return byte[]
	content, _ := ioutil.ReadFile(name.(string))
	return nextfunc(content, scan)
}

func filterCharsAndNormalize(text interface{}, nextfunc NextFunc) interface{} {
	// text: []byte
	// return []byte
	regp, _ := regexp.Compile(`[\W_]+`)
	content := text.([]byte)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	return nextfunc(strings.ToLower(string(content)), removeStopWords)
}

func scan(text interface{}, nextfunc NextFunc) interface{} {
	// text: string
	// return []string
	s := text.(string)
	return nextfunc(strings.Fields(s), frequency)
}

func removeStopWords(texts interface{}, nextfunc NextFunc) interface{} {
	// text: []string
	// return: []string
	content, _ := ioutil.ReadFile("../stop_words.txt")
	words := strings.Split(strings.TrimSpace(string(content)), ",")
	stopwords := map[string]bool{}
	for _, w := range words {
		stopwords[w] = true
	}
	res := []string{}
	for _, w := range texts.([]string) {
		if _, ok := stopwords[w]; !ok && len(w) > 1 {
			res = append(res, w)
		}
	}
	return nextfunc(res, sortPair)
}

func frequency(words interface{}, nextfunc NextFunc) interface{} {
	// words: []string
	// return map[string]int
	res := map[string]int{}
	for _, w := range words.([]string) {
		if _, ok := res[w]; ok {
			res[w]++
		} else {
			res[w] = 1
		}
	}
	return nextfunc(res, top25Str)
}

type pair struct {
	str   string
	count int
}
type pairSlice []pair

func (a pairSlice) Len() int {
	return len(a)
}
func (a pairSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a pairSlice) Less(i, j int) bool {
	return a[j].count < a[i].count
}

func sortPair(wordcounts interface{}, nextfunc NextFunc) interface{} {
	// wordcounts: map[string]int
	// return pairSlice
	countArr := pairSlice{}
	for k, v := range wordcounts.(map[string]int) {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)
	return nextfunc(countArr, printText)
}

func top25Str(pairi interface{}, nextfunc NextFunc) interface{} {
	// pairi: pairSlice
	// return string
	pairs := pairi.(pairSlice)
	res := []byte{}
	for i := 0; i < 25; i++ {
		res = append(res, []byte(pairs[i].str+" - "+strconv.Itoa(pairs[i].count)+"\n")...)
	}
	return nextfunc(string(res), nil)
}
func printText(text interface{}, nextfunc NextFunc) interface{} {
	fmt.Print(text)
	return nil
}
func main() {
	readFile(os.Args[1], filterCharsAndNormalize)
}
