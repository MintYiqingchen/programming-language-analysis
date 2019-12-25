#!/usr/bin/env python
import re, sys, operator

# Mileage may vary. If this crashes, make it lower
RECURSION_LIMIT = 50
# We add a few more, because, contrary to the name,
# this doesn't just rule recursion: it rules the 
# depth of the call stack
sys.setrecursionlimit(RECURSION_LIMIT+80)
# params[0]: word_list, params[1]: stopwords, params[2]: wordfreqs
Y = (lambda h: lambda F: F(lambda x: h(h)(F)(x)))(lambda h: lambda F: F(lambda x: h(h)(F)(x)))
count = Y(lambda f: lambda params: None if len(params[0]) == 0 else \
            (addToCount(params[0], params[1], params[2]), f((params[0][1:], params[1], params[2]))) )
wf_print = Y(lambda f: lambda wordfreq: None if len(wordfreq) == 0 else \
    (print(wordfreq[0][0], ' - ', wordfreq[0][1]), f(wordfreq[1:]))
)

def addToCount(word_list, stopwords, wordfreqs):
    word = word_list[0]
    if word not in stopwords:
        if word in wordfreqs:
            wordfreqs[word] += 1
        else:
            wordfreqs[word] = 1
    # print(len(word_list))

stop_words = set(open('../stop_words.txt').read().split(','))
words = re.findall('[a-z]{2,}', open(sys.argv[1]).read().lower())
word_freqs = {}
# Theoretically, we would just call count(words, stop_words, word_freqs)
# Try doing that and see what happens.
for i in range(0, len(words), RECURSION_LIMIT):
    count((words[i:i+RECURSION_LIMIT], stop_words, word_freqs))
wf_print(sorted(word_freqs.items(), key=operator.itemgetter(1), reverse=True)[:25])
