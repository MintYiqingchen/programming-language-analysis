import os, sys, re, math, operator
import numpy as np
from keras.models import Sequential
from keras.utils import to_categorical
import keras.layers as layers
#from keras.utils import plot_model
from keras import backend as K
from collections import Counter
stopwords = set(open('../stop_words.txt').read().split(','))
all_words = re.findall('[a-z]{2,}', open(sys.argv[1]).read().lower())
# build up vocabulary dictionary
uniqs = [''] + list(set(all_words))
uniqs_indices = dict((w, i) for i, w in enumerate(uniqs))
indices_uniqs = dict((i, w) for i, w in enumerate(uniqs))
indices = [uniqs_indices[w] for w in all_words]

WORDS_SIZE = len(all_words)
VOCAB_SIZE = len(uniqs)
BATCH_SIZE = 100
print(f'Words size {WORDS_SIZE}, vocab size {VOCAB_SIZE}, batch size {BATCH_SIZE}')

# map rule:
# stop words: word index --> 0
# non-stop words: word index --> word index
# Network input: one hot vector
# Network output: one hot vector
def set_weights(clayer):
    wb = []
    b = np.zeros((VOCAB_SIZE,), dtype=np.float32)
    w = np.eye(VOCAB_SIZE, dtype=np.float32)
    for stopword in stopwords:
        if stopword in uniqs_indices:
            idx = uniqs_indices[stopword]
            w[idx, idx] = 0
            w[idx, 0] = 1
    # Scale the whole thing down one order of magnitude
    #w = w * 0.1
    wb.append(w)
    wb.append(b)
    clayer.set_weights(wb)

def model_dense():
    print('Build model ...')
    model = Sequential([layers.Dense(VOCAB_SIZE, input_dim=VOCAB_SIZE, trainable=False)])
    set_weights(model.layers[0])
    return model, "stopwords-nolearning-{}v-{}f".format(VOCAB_SIZE, VOCAB_SIZE)

model, name = model_dense()
model.summary()

counter = Counter()
for i in range(0, len(indices), BATCH_SIZE):
    batch = indices[i: i + BATCH_SIZE]
    batch_x = to_categorical(batch, VOCAB_SIZE)
    preds = model.predict_on_batch(batch_x) # [BATCH_SIZE, VOCAB_SIZE]
    pred_indices = np.argmax(preds, axis=1)
    counter.update((indices_uniqs[idx] for idx in pred_indices))

del counter['']
for w, c in counter.most_common(25): # print for verification
    print(w, ' - ', c)