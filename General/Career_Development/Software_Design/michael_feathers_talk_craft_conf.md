# Micheal Feathers talk: Asking, Telling and Modulrity - Choices in code

https://www.youtube.com/watch?time_continue=102&v=snoMuTDiTsM

a very procedural server class
how to me improve/refactor this
  first things first, tests
  but before that IDE auto refactorings
    extract method is a safe enough one
hunt for work
  what can be extracted
compact things to make it easier to understand
  This was in a philosophy of software design
    encapsulate complexity to make it possible to understand the system at different levels
comment a variable declaration to see where it's being used
scratch refactoring
  doing enough cuts and moves to understand things better
once you can understand the code ask the question what should this design be.
what are the responsibilities of the class
  psuedo code / java to define the interfaces
In some designs you end up with a class which is a Coordinator
  which ends up with lots of dependencies on other classes
He designs the server classes with each one having only one dependency
  decouples things nicely
  minimises dependencies
Alan Kay
  The most important thing in OO is the messages
  sync / async, it's async
Middle person
  ask for some stuff, do some processing, and tell someone about it
  like the Coordinator
make things as tell as possible.
  kind of like the actor model
  minimises dependencies
prag prog
  "Tell, don't ask"
Tell someone to do something
  they tell someone else about their result
NULL, options, exceptions
  NULL is often the easiest thing to return
  use the Null object pattern
  or use Options type
In tell, don't ask
  if something goes wrong, just don't tell
Tell above, ask below
  Tell stuff seems to be at the top level
  ask / functional stuff seems to be at the bottom
  async at the highest level gives major decoupling
  down low, lazy evaluation
void return is tell
if there is a return value it's ask
pure functions are used when asking, functions with side-effects are used when telling
