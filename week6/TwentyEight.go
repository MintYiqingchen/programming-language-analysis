package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func DataStorageActor(filepath string, sendBox chan string) {
	content, _ := ioutil.ReadFile(filepath)
	regp, _ := regexp.Compile(`[\W_]+`)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	text := strings.ToLower(string(content))
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		sendBox <- words[i]
	}
	close(sendBox)
}

func StopWordActor(recvBox, sendBox chan string) {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	text := strings.Split(strings.TrimSpace(string(content)), ",")
	stopwords := map[string]bool{}
	for i := 0; i < len(text); i++ {
		stopwords[text[i]] = true
	}
	for w := range recvBox {
		if _, ok := stopwords[w]; !ok && len(w) > 1 {
			sendBox <- w
		}
	}
	close(sendBox)
}

type item struct {
	word  string
	count int
}

func FrequencyActor(recvBox chan string, sendBox chan []item) {
	counter := map[string]int{}
	for w := range recvBox {
		if _, ok := counter[w]; ok {
			counter[w]++
		} else {
			counter[w] = 1
		}
	}
	// to slice
	itemslice := []item{}
	for k, v := range counter {
		itemslice = append(itemslice, item{k, v})
	}
	// sort
	sort.Slice(itemslice, func(i, j int) bool { return itemslice[j].count < itemslice[i].count })
	sendBox <- itemslice[:25]
}
func main() {
	stInCh := make(chan string, 100)
	freqInCh := make(chan string, 100)
	freqOutCh := make(chan []item)
	go DataStorageActor(os.Args[1], stInCh) // thread 1
	go StopWordActor(stInCh, freqInCh)      // thread 2
	go FrequencyActor(freqInCh, freqOutCh)  // thread 3

	res := <-freqOutCh
	for _, p := range res {
		fmt.Println(p.word, " - ", p.count)
	}
}
