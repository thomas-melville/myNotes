# microservice design patterns

https://microservices.io/patterns/index.html

## Saga pattern

When a request from a user spans multiple microservices (in which some state is updated in each microservice) how do you ensure the system ends up in a consistent state?
If the request fails half way through, what do you do with the sub requests which have completed successfully?

We need transactions that span microservices.

A saga is a sequence of local transactions.
Each local transaction updates the database, and publishes a message/event to trigger the next local transaction.
Each local transaction must have a compensating action which is performed for a previous transactions in the saga if one fails.
  they undo the changes which were made.

A saga seems to start from an initial microservice which in the end will need to be updated also.
So, there will be a need for temporary states of objects in the service as the saga is executing

Two ways of coordinating a saga:
1. Choreography
2. orchestration

### Choreography

each local transaction publishes domain events that trigger local transactions in other services.
The other services can then publish events to trigger transactions in other services (including the original service)

### Orchestration

An orchestrator (object) tells the participants what local transactions to execute.
A single service would publish all domain events through the course of the saga, and react to events from individual services

### Issues to address when using this pattern

1. Ensure atomicity between writing to the database and sending the event
    Solution is one of a number of other microservice design patterns.
    1. Event sourcing
    2. Transactional Outbox
    3. Aggregates
    4. Domain Events

2. Informing user when saga completes.
    The initial microservice receives a synchronouns request, which triggers an asynchronous saga.
    1. don't send back the response until the Saga completes.
    2. send back an order id and let the user poll
    3. send an event, websocket/hook, to the client
