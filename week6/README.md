### Style Actor (TwentyEight.go)
Constraints:

- The larger problem is decomposed into 'things' that make sense for
  the problem domain 

- Each 'thing' is a capsule of data that exposes one single procedure,
  namely the ability to receive and dispatch messages that are sent to
  it

- Message dispatch can result in sending the message to another capsule

Possible names:

- Letterbox
- Messaging style
- Objects
- Actors

When receive a message, a actor can:
* change its state
* send more message
* create new actors

### Style Dataspace (TwentyNine.go)
Constraints:

- Existence of one or more units that execute concurrently

- Existence of one or more data spaces where concurrent units store and
  retrieve data

- No direct data exchanges between the concurrent units, other than via the data spaces

Possible names:

- Dataspaces
- Linda

### Style double MapReduce (ThirtyOne.js)
Constraints:

- Input data is divided in chunks, similar to what an inverse multiplexer does to input signals

- A map function applies a given worker function to each chunk of data, potentially in parallel

- The results of the many worker functions are reshuffled in a way
  that allows for the reduce step to be also parallelized

- The reshuffled chunks of data are given as input to a second map
  function that takes a reducible function as input

Possible names:

- Map-reduce 
- Hadoop style
- Double inverse multiplexer 
