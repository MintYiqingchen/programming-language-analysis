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

type TFTheOne struct {
	val interface{}
}

func (obj *TFTheOne) bind(myfunc func(interface{}) interface{}) *TFTheOne {
	obj.val = myfunc(obj.val)
	return obj
}

func (obj *TFTheOne) printme() {
	fmt.Print(obj.val)
}
func readFile(name interface{}) interface{} {
	// name : string
	// return byte[]
	content, _ := ioutil.ReadFile(name.(string))
	return interface{}(content)
}

func filterCharsAndNormalize(text interface{}) interface{} {
	// text: []byte
	// return []byte
	regp, _ := regexp.Compile(`[\W_]+`)
	content := text.([]byte)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	return strings.ToLower(string(content))
}

func scan(text interface{}) interface{} {
	// text: string
	// return []string
	s := text.(string)
	return strings.Fields(s)
}

func removeStopWords(texts interface{}) interface{} {
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
	return res
}

func frequency(words interface{}) interface{} {
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
	return res
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

func sortPair(wordcounts interface{}) interface{} {
	// wordcounts: map[string]int
	// return pairSlice
	countArr := pairSlice{}
	for k, v := range wordcounts.(map[string]int) {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)
	return countArr
}

func top25Str(pairi interface{}) interface{} {
	// pairi: pairSlice
	// return string
	pairs := pairi.(pairSlice)
	res := []byte{}
	for i := 0; i < 25; i++ {
		res = append(res, []byte(pairs[i].str+" - "+strconv.Itoa(pairs[i].count)+"\n")...)
	}
	return string(res)
}
func main() {
	obj := TFTheOne{os.Args[1]}
	obj.bind(readFile).bind(filterCharsAndNormalize).bind(scan).bind(removeStopWords).bind(frequency).bind(sortPair).bind(top25Str).printme()
}
