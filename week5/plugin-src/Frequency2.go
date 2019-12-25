package main

import "container/heap"

type Pair struct {
	Str   string
	Count int
}

func (p Pair) GetStr() string {
	return p.Str
}
func (p Pair) GetCount() int {
	return p.Count
}

type myHeap []Pair

func (h myHeap) Len() int           { return len(h) }
func (h myHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h myHeap) Less(i, j int) bool { return h[j].Count < h[i].Count }
func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *myHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
func Top25(wordlist []string) interface{} {
	wordcounter := map[string]int{}
	for _, w := range wordlist {
		if _, ok := wordcounter[w]; ok {
			wordcounter[w]++
		} else {
			wordcounter[w] = 1
		}
	}
	pairs := myHeap{}
	for k, v := range wordcounter {
		pairs = append(pairs, Pair{k, v})
	}
	heap.Init(&pairs)
	res := make([]interface{}, 0)
	for i := 0; i < 25 ; i++ {
	    res = append(res, heap.Pop(&pairs))
	}
	return res
}
