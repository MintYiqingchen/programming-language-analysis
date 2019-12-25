import os, sys, re, math, operator
import numpy as np
from keras.models import Sequential
from keras.utils import to_categorical
import keras.layers as layers
from keras.optimizers import SGD
from keras import backend as K
from collections import Counter
stopwords = set(open('../stop_words.txt').read().split(','))
all_words = re.findall('[a-z]{2,}', open(sys.argv[1]).read().lower())
words = [w for w in all_words if w not in stopwords]
# build up vocabulary dictionary
vocabulary = [''] + list(set(words)) + list(stopwords)
posi_vocabulary = vocabulary[:-len(stopwords)]

word_indices = dict((w, i) for i, w in enumerate(vocabulary))
indices_word = dict((i, w) for i, w in enumerate(vocabulary))
indices = [word_indices[w] for w in all_words] # test set

NON_STOP_NUM = len(posi_vocabulary)
VOCAB_SIZE = len(vocabulary)
BATCH_SIZE = 100
EPOCH_NUM = 3000
EMBED_DIM = 100
PROPORTION = NON_STOP_NUM / VOCAB_SIZE
POS_BATCH_SIZE = int(BATCH_SIZE * PROPORTION)
NEG_BATCH_SIZE = BATCH_SIZE - POS_BATCH_SIZE
print(f'Nonstop Words size {NON_STOP_NUM}, vocab size {VOCAB_SIZE}, batch size {POS_BATCH_SIZE} {NEG_BATCH_SIZE}')

def generate_trainset():
    pos_x = np.random.randint(NON_STOP_NUM, size=POS_BATCH_SIZE)
    pos_y = pos_x[:]
    neg_x = np.random.randint(NON_STOP_NUM, VOCAB_SIZE, size=NEG_BATCH_SIZE)
    neg_y = np.zeros(len(neg_x))
    x = np.concatenate((pos_x, neg_x), axis=0).astype(np.float32)
    y = np.concatenate((pos_y, neg_y), axis=0).astype(np.float32)
    return x, y

def model_dense():
    print('Build model ...')
    model = Sequential([
        layers.Embedding(VOCAB_SIZE, EMBED_DIM),
        layers.Reshape((EMBED_DIM,)),
        layers.Dense(VOCAB_SIZE, activation="softmax")
        ])
    for l in model.layers:
        print(l.output_shape)
    return model

model = model_dense()
sgd = SGD(lr=1, decay=1e-6, momentum=0.9)
model.compile(loss='categorical_crossentropy', optimizer=sgd, metrics=['accuracy'])
model.summary()

# -------------- Training -------------------
for e in range(EPOCH_NUM):
    batch_x, batch_y = generate_trainset()
    batch_x = np.array(batch_x)[:, np.newaxis]
    batch_y = to_categorical(batch_y, VOCAB_SIZE)
    metrics = model.train_on_batch(batch_x, batch_y)
    if e % 200 == 0:
        print(e, model.metrics_names, metrics)

# --------------- Testing --------------------
counter = Counter()
for i in range(0, len(indices), BATCH_SIZE):
    batch = indices[i: i + BATCH_SIZE]
    batch_x = np.array(batch)[:, np.newaxis]
    preds = model.predict_on_batch(batch_x) # [BATCH_SIZE, VOCAB_SIZE]
    pred_indices = np.argmax(preds, axis=1)
    counter.update((indices_word[idx] for idx in pred_indices))

del counter['']
for w, c in counter.most_common(25): # print for verification
    print(w, ' - ', c)