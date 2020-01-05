### Style Pub-sub (Fifteen.go)

Constraints:

- Larger problem is decomposed into entities using some form of abstraction
  (objects, modules or similar)

- The entities are never called on directly for actions

- Existence of an infrastructure for publishing and subscribing to
  events, aka the bulletin board

- Entities post event subscriptions (aka 'wanted') to the bulletin
  board and publish events (aka 'offered') to the bulletin board. the
  bulletin board does all the event management and distribution

Possible names:

- Bulletin board
- Publish-Subscribe

### Style Prototype (Twelve.js)


Constraints:

- The larger problem is decomposed into 'things' that make sense for
  the problem domain 

- Each 'thing' is a map from keys to values. Some values
are procedures/functions.

Possible names:

- Closed Maps
- Prototypes
