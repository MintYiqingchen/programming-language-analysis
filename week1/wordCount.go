package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

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

func loadStopWords() map[string]bool {
	fp, err := os.Open("../stop_words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	s, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(s), ",")
	res := map[string]bool{}
	for _, w := range words {
		res[w] = true
	}
	tmp := []byte{'a'}
	for tmp[0] <= 'z' {
		res[string(tmp)] = true
		tmp[0]++

	}
	return res
}
func normalize(line string) string {
	line = strings.ToLower(line)
	linesl := []byte(line)
	for i := 0; i < len(linesl); i++ {
		if linesl[i] <= byte('z') && linesl[i] >= byte('a') {
			continue
		}
		linesl[i] = ' '
	}
	return string(linesl)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: wordCount text_file_name")
		return
	}
	// read stop_words.txt
	stopWords := loadStopWords()

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// read every line
	reader := bufio.NewReader(fp)
	wordCounter := map[string]int{}
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if s == "\n" {
			continue
		}
		s = normalize(s)
		fields := strings.Fields(s)
		for _, field := range fields {
			if _, ok := stopWords[field]; ok {
				continue
			}
			if _, ok := wordCounter[field]; ok {
				wordCounter[field]++
			} else {
				wordCounter[field] = 1
			}
		}
	}

	// sort by count
	countArr := pairSlice{}
	for k, v := range wordCounter {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)

	for i := 0; i < 25; i++ {
		fmt.Println(countArr[i].str, " - ", countArr[i].count)
	}
}
