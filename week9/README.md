### Style Neural network
1. easy.py
Change count_words_no_learning (Links to an external site.) so that it doesn't crash while processing pride-and-prejudice.txt. You will likely need to predict in smaller batches than the entire file, and hold an external dictionary in the Python program to accumulate the counts. (NNs don't have memory)  
2. moderate0.py
Program a DNN for eliminating stop words from given lines. That is, set up the model and hard-code the weights.  Input: a sequence of words coded either as one-hot or as integers. Output: another sequence of words, with stop words eliminated. 
3. moderate1.py
Train a DNN for eliminating stop words from given lines. Input: a sequence of words coded either as one-hot or as integers. Output: another sequence of words, with stop words eliminated. 

### Run
My version:
python3, tensorflow==1.12.0, keras==2.2.4

RUN:
```
python easy.py ../pride-and-prejudice.txt
python moderate0.py ../pride-and-prejudice.txt
python moderate1.py ../pride-and-prejudice.txt
```