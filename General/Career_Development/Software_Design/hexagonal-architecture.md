# hexagonal architecture

https://fideloper.com/hexagonal-architecture
https://www.youtube.com/watch?v=AOIWUPjal60
https://jmgarridopaz.github.io/

Instead of layers from top to bottom the architecture is a hexagonal shape with your core business and domain logic are at the core
There are layers around this for communicating with external systems and supporting tools.
There is the driving side, and the driven side.
  driving side, what drives the application.
  driven side, what the application drives.

Layers of code responsibility, decoupling between layers

decoupling of code from our framework
letting our application express itself
using a framework as a means to accomplish tasks in our application
instead of being our application itself

Alistar Cockburn on it's intent:

"Allow an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases."

The number of sides is arbitrary, the point is it has many sides.
each side represents a "port" into/out of our application

a port is a way to accept data into our application
"An access port for a kind of a protocol" Two way communication!
Ports collect a kind of a protocol, with requests and responses.
any protocol is a port, Rest, message, etc.
connecting to a database is a port

Maintainability
Low Technical Debt

"Identify the aspects that VARY and separate them from what stays the SAME"

Each layer has two elements:

* Code
  actual code doing things
  often this is code which acts as adapters to ports
  busines logic, ...
* Boundary
  the ports
  how outside layers can communicate to the current layer

Layers:

* Domain
  inner most layer
  your business logic
  and defines how the layer outside interacts with it.
  policy, rules your code must follow
  it's what gives your application value
  Domain Driven Design (Distilled)
* Application
  orchestrates the use of the entities found in the Domain layer
  adapts requests from the framework
* Framework
  contains code your application uses but it is not actually your application
  often literally your framework, Spring.
  but can be any other libraries too
  Implements services defined by the application layer
  also adapts requests from outside our application
    http requests, gather inputs and route it to controller
Outside

The hexagonal architecture is also known as the Ports and Adapters architecture

## Ports and Adapters

ports which can be adapted for any given layer
Layers communicate with each other using interfaces (ports) and implementations (adapters)

Inside to Outside

Interfaces, defines a contract.
Any user of the interface doesn't care about the implementation.
interface stays the same while the implementation can vary.

A way to add cross cutting concerns to all interface implementations is to use the decorator pattern

## Boundaries

Each layer defines how the next outside layer communicates with it.
Using an Interface.
At each layer boundary we find interfaces.
These interfaces are the ports for the next layer to create adapters (implementations) for

The Domain layer defines use cases, it's job is just to say "this is how you can use me"
The application layer is responsible for orchestrating the Domain layer code in order to accomplish a task.
This repeats as we go out each layer.
Adapters (Implementations) for the inside layers ports, and it's own ports for the next outside layer to implement

## Use cases / commands

ports/adapters are on a micro level with the software.
Application Boundary is a macro level concept.
This boundary separates our application as a whole from everything else.
Use cases/command are strict definitions of how the outside world can communicate with our application.
Essentially classes which name actions that can be taken
*****

## Designing process

1. Identify actors
2. Identify ports
3. Add adapters
4. The whole picture
5. Driver ports

This would result in a diagram showing the components of the system: Hexagon (with ports) actors and adapters.
Driver ports in detail.
Driven ports not defined in detail yet.

### Identify actors

1. Draw a hexagon
2. Put driver actors outside the hexagon, on the left and upper side.
3. Pur driven actors outside the hexagon, on the right and lower side.

### Identify ports

#### Driver ports

For each driver port, ask yourself: what purpose does the driver actor want the app for?
  The answer will be the name of the port.
  for <action verb> <noun>

### Add Adapters

For each port create two adapters:
1. Mock adapter
    use this in testing
      on the driver side it's the test automation tool
      on the driven side it's a mock
2. Real adapter
    use this in production
