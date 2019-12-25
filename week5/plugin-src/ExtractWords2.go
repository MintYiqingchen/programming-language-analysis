package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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
func ExtractWords(filename string) []string {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	res := make([]string, 0)
	stopWords := loadStopWords()
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
			if _, ok := stopWords[field]; !ok {
				res = append(res, field)
			}
		}
	}
	return res
}
