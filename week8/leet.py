#!/usr/bin/python3
import numpy as np
import sys
from collections import Counter
import re
leet_map = {
    'a':'4','b':'13','c':'[','d':'|)','e':'3','f':'ph',
    'g':'6','h':'#','i':'1','j':']','k':'|<','l':'1',
    'm':'/\/\\','n':'|\|','o':'0','p':'|>','q':'O_','r':'I2',
    's':'5','t':'7','u':'(_)','v':'\/','w':'\/\/','x':'><',
    'y':'j','z':'2',
    '0':'()','1':'I','2':'Z','3':'E','4':'A','5':'S',
    '6':'b','7':'T','8':'B','9':'q'
}

def extract_words(path_to_file):
    with open(path_to_file) as f:
        str_data = f.read()
    pattern = re.compile('[\W_]+')
    word_list = pattern.sub(' ', str_data).lower().split()
    return np.array(word_list, dtype=np.string_) # string array

def strToLeetSpeak(npstr):
    tmp = npstr.tostring().decode("utf8")
    tmp = "".join(map(lambda ch: leet_map[ch] if ch in leet_map else ch, tmp))
    #print(tmp)
    return tmp

def make2gram(strary):
    nextelem = strary[1:]
    strary = strary[:-1]
    return list(map(lambda x :x[0]+" "+x[1], zip(strary, nextelem)))

words = extract_words(sys.argv[1])
words = list(map(strToLeetSpeak, words))
words = make2gram(words)
counter = Counter(words)
list(map(lambda x: print(x[0], ' - ', x[1]), counter.most_common(5)))