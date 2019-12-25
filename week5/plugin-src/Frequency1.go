package main

import "sort"

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
func Top25(wordlist []string) interface{} {
	wordcounter := map[string]int{}
	for _, w := range wordlist {
		if _, ok := wordcounter[w]; ok {
			wordcounter[w]++
		} else {
			wordcounter[w] = 1
		}
	}
	pairs := []Pair{}
	for k, v := range wordcounter {
		pairs = append(pairs, Pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[j].Count < pairs[i].Count })
	res := make([]interface{}, 0)
	for i := range pairs[:25] {
		res = append(res, pairs[i])
	}
	return res
}
