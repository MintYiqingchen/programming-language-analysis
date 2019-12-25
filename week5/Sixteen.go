package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func readStopWords() map[string]bool {
	//  """ This function can only be called from a function named removeStopWords."""
	if pc, _, _, _ := runtime.Caller(1); runtime.FuncForPC(pc).Name() != "main.(*MainObject).RemoveStopWords" {
		return nil
	}
	// pc, _, _, _ := runtime.Caller(1)
	// fmt.Println(runtime.FuncForPC(pc).Name())
	content, _ := ioutil.ReadFile("../stop_words.txt")
	words := strings.Split(strings.TrimSpace(string(content)), ",")
	res := map[string]bool{}
	for _, w := range words {
		res[w] = true
	}
	return res
}

type MainObject struct{}

func (obj *MainObject) ReadFile(name string) string {
	content, _ := ioutil.ReadFile(name)
	regp, _ := regexp.Compile(`[\W_]+`)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	return strings.ToLower(string(content))
}

func (obj *MainObject) Scan(text string) []string {
	return strings.Fields(text)
}

func (obj *MainObject) RemoveStopWords(texts []string) []string {
	stopwords := readStopWords()
	res := []string{}
	for _, w := range texts {
		if _, ok := stopwords[w]; !ok && len(w) > 1 {
			res = append(res, w)
		}
	}
	return res
}

func (obj *MainObject) Frequency(words []string) map[string]int {
	res := map[string]int{}
	for _, w := range words {
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
func (obj *MainObject) SortPair(wordcounts map[string]int) pairSlice {
	countArr := pairSlice{}
	for k, v := range wordcounts {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)
	return countArr
}

// reflection helper function
func CallByName(obj *MainObject, funcname string, params ...interface{}) ([]interface{}, error) {
	fv := reflect.ValueOf(obj).MethodByName(funcname)
	if !fv.IsValid() {
		fmt.Println(funcname)
		return nil, fmt.Errorf("method %s not found", funcname)
	}
	callin := make([]reflect.Value, len(params))
	for i := range params {
		callin[i] = reflect.ValueOf(params[i])
	}
	ret := make([]interface{}, 0)
	for _, v := range fv.Call(callin) {
		ret = append(ret, v.Interface())
	}

	return ret, nil
}
func main() {
	obj := &MainObject{}
	ret, _ := CallByName(obj, "ReadFile", os.Args[1])
	ret, _ = CallByName(obj, "Scan", ret...)
	ret, _ = CallByName(obj, "RemoveStopWords", ret...)
	ret, _ = CallByName(obj, "Frequency", ret...)
	ret, _ = CallByName(obj, "SortPair", ret...)
	wordcount := ret[0].(pairSlice)
	for _, val := range wordcount[:25] {
		fmt.Println(val.str, " - ", val.count)
	}
}
