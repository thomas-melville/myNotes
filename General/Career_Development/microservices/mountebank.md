# mountbank

http://www.mbtest.org

create test doubles which are accessed **over the network.**
looks to be a node js module

you post the responses you want to give that match certain predicates on a specific port.
You are creating the "imposter"

other features include:

* record-replay
* fault injection

The specific port is then opened and when you send a request it figures out which response to send based on the request

works on http and tcp (which helps when you're talking RPC)

You then need to configure your service to talk to the imposter

## terms

* imposter
* * a server representing a test double
* * identified by port and protocol on the host that mountebank is running
* stub
* * A set of configuration used to generate a response for an imposter.
* * imposter can have 0 or more stubs
* * each of which are associated with different predicates.
* predicate
* * A condition that determines whether a stub is responsible for responding.
* * each stub can have 0 or more predicates
