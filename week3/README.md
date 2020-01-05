### Style Y-combinator (seven.py)

Constraints:

- All, or a significant part, of the problem is modelled by
  induction. That is, specify the base case (n_0) and then the n+1
  rule

Possible names:

- Infinite mirror
- Inductive
- Recursive

### Style Kick forward (Eight.go)

Variation of the candy factory style, with the following additional constraints:

- Each function takes an additional parameter, usually the last, which is another function
- That function parameter is applied at the end of the current function
- That function parameter is given as input what would be the output of the current function
- Larger problem is solved as a pipeline of functions, but where the next function to be applied is given as parameter to the current function

Possible names:

- Kick your teammate forward!
- Continuation-passing style
- Crochet loop

### Style The one (Nine.go)

Constraints:

- Existence of an abstraction to which values can be
converted. 

- This abstraction provides operations to (1) wrap
around values, so that they become the abstraction; (2) bind
itself to functions, so to establish sequences of functions;
and (3) unwrap the value, so to examine the final result.

- Larger problem is solved as a pipeline of functions bound
together, with unwrapping happening at the end.

- Particularly for The One style, the bind operation simply
calls the given function, giving it the value that it holds, and holds
on to the returned value.


Possible names:

- The One
- Monadic Identity
- The wrapper of all things
- Imperative functional style