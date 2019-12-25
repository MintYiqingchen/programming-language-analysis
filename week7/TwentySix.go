package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func loadAllWords() interface{} {
	regp := regexp.MustCompile("[a-z]{2,}")
	content, _ := ioutil.ReadFile(os.Args[1])
	words := regp.FindAllString(strings.ToLower(string(content)), -1)
	return words
}

func loadStopWords() interface{} {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	text := strings.Split(strings.TrimSpace(string(content)), ",")
	stopwords := map[string]bool{}
	for i := 0; i < len(text); i++ {
		stopwords[text[i]] = true
	}
	return stopwords
}

type Column struct {
	Value interface{}
	Func  func() interface{}
}
type Pair struct {
	word  string
	count int
}

var allWords, stopWords, nonStopWords, uniqueWords, counts, sortedData Column
var allColumns []*Column

func update() {
	for i := range allColumns {
		if allColumns[i].Func != nil {
			allColumns[i].Value = allColumns[i].Func()
		}
	}
}
func nonStopWordsFunc() interface{} {
	res := []string{}
	words := allWords.Value.([]string)
	stopwords := stopWords.Value.(map[string]bool)
	for i := range words {
		if _, ok := stopwords[words[i]]; ok || len(words[i]) <= 1 {
			res = append(res, "")
		} else {
			res = append(res, words[i])
		}
	}
	return res
}
func uniqueWordsFunc() interface{} {
	res := []string{}
	set := map[string]bool{}
	for _, w := range nonStopWords.Value.([]string) {
		if _, ok := set[w]; !ok && w != "" {
			res = append(res, w)
			set[w] = true
		}
	}
	return res
}
func countsFunc() interface{} {
	res := []int{}
	uniquewords := uniqueWords.Value.([]string)
	countsMap := map[string]int{}
	for i := range uniquewords {
		countsMap[uniquewords[i]] = 0
	}
	for _, w := range nonStopWords.Value.([]string) {
		if _, ok := countsMap[w]; ok {
			countsMap[w]++
		}
	}
	for i := range uniquewords {
		res = append(res, countsMap[uniquewords[i]])
	}
	return res
}
func sortedDataFunc() interface{} {
	res := []Pair{}
	words := uniqueWords.Value.([]string)
	nums := counts.Value.([]int)
	for i := range words {
		res = append(res, Pair{words[i], nums[i]})
	}
	sort.Slice(res, func(i, j int) bool { return res[i].count > res[j].count })
	return res
}
func main() {
	allWords = Column{loadAllWords(), nil}
	stopWords = Column{loadStopWords(), nil}
	nonStopWords = Column{nil, nonStopWordsFunc}
	uniqueWords = Column{nil, uniqueWordsFunc}
	counts = Column{nil, countsFunc}
	sortedData = Column{nil, sortedDataFunc}
	allColumns = []*Column{&allWords, &stopWords, &nonStopWords, &uniqueWords, &counts, &sortedData}
	update()
	for _, p := range sortedData.Value.([]Pair)[:25] {
		fmt.Println(p.word, " - ", p.count)
	}
}
