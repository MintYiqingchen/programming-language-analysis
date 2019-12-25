package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

var stopwords map[string]bool = map[string]bool{}

func loadStopWords() {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	text := strings.Split(strings.TrimSpace(string(content)), ",")
	for i := 0; i < len(text); i++ {
		stopwords[text[i]] = true
	}
}

type threadsafeQueue struct {
	array []interface{}
	mx    sync.Mutex
}

func (q *threadsafeQueue) Put(v interface{}) {
	q.mx.Lock()
	defer q.mx.Unlock()
	q.array = append(q.array, v)
}
func (q *threadsafeQueue) Get() (interface{}, bool) {
	q.mx.Lock()
	defer q.mx.Unlock()
	if len(q.array) == 0 {
		return nil, false
	}
	res := q.array[0]
	q.array = q.array[1:]
	return res, true
}

var wordspace threadsafeQueue
var alphabetspace [26]threadsafeQueue
var freqspace threadsafeQueue

func processWords(syncCh chan struct{}) {
	var wordfreqs [26]map[string]int
	for i := 0; i < 26; i++ {
		wordfreqs[i] = make(map[string]int)
	}
	for {
		w, ok := wordspace.Get()
		if !ok {
			break
		}
		word := w.(string)
		if _, ok := stopwords[word]; ok || len(word) <= 1 {
			continue
		}
		idx := int(word[0] - 'a')
		if _, ok := wordfreqs[idx][word]; ok {
			wordfreqs[idx][word]++
		} else {
			wordfreqs[idx][word] = 1
		}
	}
	for i := 0; i < 26; i++ {
		alphabetspace[i].Put(wordfreqs[i])
	}
	syncCh <- struct{}{}
}

func mergeCount(id int, syncCh chan struct{}) {
	wordfreqs := map[string]int{}
	for id < 26 {
		for i := 0; i < len(alphabetspace[id].array); i++ {
			dict := alphabetspace[id].array[i].(map[string]int)
			for k, v := range dict {
				if _, ok := wordfreqs[k]; ok {
					wordfreqs[k] += v
				} else {
					wordfreqs[k] = v
				}
			}
		}
		id += 6
	}
	freqspace.Put(wordfreqs)
	syncCh <- struct{}{}
}
func main() {
	loadStopWords()
	regp := regexp.MustCompile("[a-z]{2,}")
	content, _ := ioutil.ReadFile(os.Args[1])
	words := regp.FindAllString(strings.ToLower(string(content)), -1)
	for i := range words {
		wordspace.Put(words[i])
	}
	syncCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		go processWords(syncCh)
	}
	for i := 0; i < 5; i++ {
		<-syncCh
	}
	// 29.2
	for i := 0; i < 5; i++ {
		go mergeCount(i, syncCh)
	}
	for i := 0; i < 5; i++ {
		<-syncCh
	}
	type item struct {
		word  string
		count int
	}
	pairs := []item{}
	for i := 0; i < len(freqspace.array); i++ {
		wordfreqs := freqspace.array[i].(map[string]int)
		for k, v := range wordfreqs {
			pairs = append(pairs, item{k, v})
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })
	for i := 0; i < 25; i++ {
		fmt.Println(pairs[i].word, " - ", pairs[i].count)
	}
}
