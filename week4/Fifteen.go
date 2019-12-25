package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Event struct {
	etype string
	param interface{}
}
type Handler func(*Event)
type HandlerSlice []Handler
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

// EventManager class
type EventManager struct {
	subscriptions map[string]HandlerSlice
}

func (em *EventManager) subscribe(etype string, handle Handler) {
	if _, ok := em.subscriptions[etype]; ok {
		em.subscriptions[etype] = append(em.subscriptions[etype], handle)
	} else {
		em.subscriptions[etype] = HandlerSlice{handle}
	}
}

func (em *EventManager) publish(event *Event) {
	tevent := event.etype
	if _, ok := em.subscriptions[tevent]; ok {
		for i := 0; i < len(em.subscriptions[tevent]); i++ {
			em.subscriptions[tevent][i](event)
		}
	}
}

// DataStorage class
type DataStorage struct {
	em    *EventManager
	_data string
}

// - constructor
func NewDataStorage(em *EventManager) *DataStorage {
	res := &DataStorage{em, ""}
	em.subscribe("load", res.Load)
	em.subscribe("start", res.ProduceWords)
	return res
}
func (ds *DataStorage) Load(event *Event) {
	filepath := event.param.(string)
	content, _ := ioutil.ReadFile(filepath)
	regp, _ := regexp.Compile(`[\W_]+`)
	content = regp.ReplaceAllLiteral(content, []byte(" "))
	ds._data = strings.ToLower(string(content))
}
func (ds *DataStorage) ProduceWords(event *Event) {
	data := strings.Fields(ds._data)
	for _, w := range data {
		ds.em.publish(&Event{"word", w})
	}
	ds.em.publish(&Event{"eof", nil})
}

// StopWordFilter class
type StopWordFilter struct {
	em        *EventManager
	stopwords map[string]bool
}

func NewStopWordFilter(em *EventManager) *StopWordFilter {
	res := &StopWordFilter{em, map[string]bool{}}
	em.subscribe("load", res.Load)
	em.subscribe("word", res.IsStopWord)
	return res
}
func (filter *StopWordFilter) Load(event *Event) {
	content, _ := ioutil.ReadFile("../stop_words.txt")
	words := strings.Split(strings.TrimSpace(string(content)), ",")
	for _, w := range words {
		filter.stopwords[w] = true
	}
}
func (filter *StopWordFilter) IsStopWord(event *Event) {
	word := event.param.(string)
	if _, ok := filter.stopwords[word]; len(word) > 1 && !ok {
		filter.em.publish(&Event{"valid_word", word})
	}
}

// WordFrequencyCounter
type WordFrequencyCounter struct {
	em        *EventManager
	wordFreqs map[string]int
}

func NewWordFrequencyCounter(em *EventManager) *WordFrequencyCounter {
	res := &WordFrequencyCounter{em, map[string]int{}}
	em.subscribe("valid_word", res.IncrementCount)
	em.subscribe("print", res.PrintFreqs)
	return res
}
func (counter *WordFrequencyCounter) IncrementCount(event *Event) {
	word := event.param.(string)
	if _, ok := counter.wordFreqs[word]; ok {
		counter.wordFreqs[word]++
	} else {
		counter.wordFreqs[word] = 1
	}
}
func (counter *WordFrequencyCounter) PrintFreqs(event *Event) {
	countArr := pairSlice{}
	for k, v := range counter.wordFreqs {
		countArr = append(countArr, pair{k, v})
	}
	sort.Sort(countArr)
	for i := 0; i < 25; i++ {
		fmt.Println(countArr[i].str, " - ", countArr[i].count)
	}
}

// WordFrequencyApplication class
type WordFrequencyApplication struct {
	em *EventManager
}

func NewWordFrequencyApplication(em *EventManager) *WordFrequencyApplication {
	res := &WordFrequencyApplication{em}
	em.subscribe("run", res.Run)
	em.subscribe("eof", res.Stop)
	return res
}
func (app *WordFrequencyApplication) Run(event *Event) {
	path := event.param.(string)
	app.em.publish(&Event{"load", path})
	app.em.publish(&Event{"start", nil})
}
func (app *WordFrequencyApplication) Stop(event *Event) {
	app.em.publish(&Event{"print", nil})
}

// 15.2
type ZWordCounter struct {
	em         *EventManager
	zwordCount int
}

func NewZWordCounter(em *EventManager) *ZWordCounter {
	res := &ZWordCounter{em, 0}
	em.subscribe("valid_word", res.IncrementZCount)
	em.subscribe("print", res.PrintCount)
	return res
}
func (counter *ZWordCounter) IncrementZCount(event *Event) {
	word := event.param.(string)
	// print all words that contains "z", duplicate words are also counted
	if strings.Contains(word, "z") {
		counter.zwordCount++
	}
}
func (counter *ZWordCounter) PrintCount(event *Event) {
	fmt.Println("15.2 task - the number of words with z: ", counter.zwordCount)
}

func main() {
	em := &EventManager{map[string]HandlerSlice{}}
	NewDataStorage(em)
	NewStopWordFilter(em)
	NewWordFrequencyCounter(em)
	NewWordFrequencyApplication(em)
	NewZWordCounter(em)
	em.publish(&Event{"run", os.Args[1]})
}
