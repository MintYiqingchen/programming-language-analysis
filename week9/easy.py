import os, sys, re, math, operator
import numpy as np
from keras.models import Sequential
import keras.layers as layers
#from keras.utils import plot_model
from keras import backend as K
from collections import defaultdict
stopwords = set(open('../stop_words.txt').read().split(','))
all_words = re.findall('[a-z]{2,}', open(sys.argv[1]).read().lower())
words = [w for w in all_words if w not in stopwords]
# build up vocabulary dictionary
uniqs = [''] + list(set(words))
uniqs_indices = dict((w, i) for i, w in enumerate(uniqs))
indices_uniqs = dict((i, w) for i, w in enumerate(uniqs))
indices = [uniqs_indices[w] for w in words]

WORDS_SIZE = len(words)
VOCAB_SIZE = len(uniqs)
BIN_SIZE = math.ceil(math.log(VOCAB_SIZE, 2))
BATCH_SIZE = 500
print(f'Words size {WORDS_SIZE}, vocab size {VOCAB_SIZE}, bin size {BIN_SIZE}')

def encode_binary(W):
    x = np.zeros((BATCH_SIZE, BIN_SIZE))
    for i, w in enumerate(W):
        for n in range(BIN_SIZE): 
            n2 = pow(2, n)
            x[i, n] = 1 if (w & n2) == n2 else 0
    return x

def set_weights(clayer):
    wb = []
    b = np.zeros((VOCAB_SIZE,), dtype=np.float32)
    w = np.zeros((BIN_SIZE, VOCAB_SIZE), dtype=np.float32)
    for i in range(VOCAB_SIZE):
        for n in range(BIN_SIZE):
            n2 = pow(2, n)
            w[n][i] = 1 if (i & n2) == n2 else -1 #-(BIN_SIZE-1)
    for i in range(VOCAB_SIZE):
        slice_1 = w[:, i]
        n_ones = len(slice_1[ slice_1 == 1 ])
        if n_ones > 0: 
            slice_1[ slice_1 == 1 ] = 1./n_ones 
        n_ones = len(slice_1[ slice_1 == -1 ])
        if n_ones > 0: 
            slice_1[ slice_1 == -1 ] = -1./n_ones 
    # Scale the whole thing down one order of magnitude
    #w = w * 0.1
    wb.append(w)
    wb.append(b)
    clayer.set_weights(wb)

def SumPoolingBatch(x):
    return K.sum(x, axis=0)
def model_dense():
    print('Build model ...')
    model = Sequential([
        layers.Dense(VOCAB_SIZE, input_dim=BIN_SIZE, trainable=False),
        layers.ReLU(threshold=1-1/BIN_SIZE),
        layers.Lambda(SumPoolingBatch)
        ])
    set_weights(model.layers[0])
    return model, "words-nolearning-{}v-{}f".format(VOCAB_SIZE, BIN_SIZE)

model, name = model_dense()
model.summary()
#plot_model(model, to_file=name + '.png', show_shapes=True)

counter = defaultdict(int)
for i in range(0, len(indices), BATCH_SIZE):
    batch = indices[i: i + BATCH_SIZE]
    batch_x = encode_binary(batch)
    preds = model.predict_on_batch(batch_x) # [VOCAB_SIZE]
    
    for idx, c in enumerate(preds):
        counter[uniqs[idx]] += c

for w, c in sorted(counter.items(), key=operator.itemgetter(1), reverse=True)[:25]:
    print(w, " - ", int(c))