package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

func readFile(name string) []string {
	fp, err := os.Open(name)
	var res []string

	if err != nil {
		fmt.Println(err)
		return res
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if s == "\n" {
			continue
		}
		res = append(res, strings.TrimSpace(s))
	}
	return res
}

func filterCharsAndNormalize(text []string) []string {
	var res = make([]string, len(text))
	copy(res, text)
	regp, _ := regexp.Compile(`[\W_]`)
	for i := range res {
		res[i] = regp.ReplaceAllLiteralString(res[i], " ")
		res[i] = strings.ToLower(res[i])
	}
	return res
}

func scan(text []string) [][]string {
	res := make([][]string, len(text))
	for i, sentence := range text {
		res[i] = strings.Fields(sentence)
	}
	return res
}

func removeStopWords(texts [][]string) []string {
	fp, _ := os.Open("../stop_words.txt")
	reader := bufio.NewReader(fp)
	s, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(s), ",")
	stopwords := map[string]bool{}
	for _, w := range words {
		stopwords[w] = true
	}
	tmp := []byte{'a'}
	for tmp[0] <= 'z' {
		stopwords[string(tmp)] = true
		tmp[0]++
	}
	res := []string{}
	for _, line := range texts {
		for _, w := range line {
			if _, ok := stopwords[w]; !ok {
				res = append(res, w)
			}
		}
	}
	return res
}

func frequency(words []string) map[string]int {
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
func sortPair(wordcounts map[string]int) pairSlice {
	countArr := pairSlice{}
	for k, v := range wordcounts {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)
	return countArr
}
func printAll(words pairSlice) {
	for _, val := range words[:25] {
		fmt.Println(val.str, " - ", val.count)
	}
}
func main() {
	printAll(sortPair(frequency(removeStopWords(scan(filterCharsAndNormalize(readFile(os.Args[1])))))))
}
