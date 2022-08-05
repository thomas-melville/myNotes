# Ports and Adapters Pattern

Also known as Hexagonal Architecture

## Intent

Allow an application to equally be driven by users, programs, automated tests or batch scripts.
And to be developed and tested in isolation from its eventual run-time devices and databases.

Define the structure of an application so that it could be run by different kinds of clients, and could be tested in isolation from external devices.

Input/event arrives from the outside at a port, a tech specific adapter converts it into a usable procedural call or message and passes it to the application.

## Nature of the solution

Don't look at the application like layers stacked on top of each other.
Look at it from the view of the inside out.
The Business logic is at the centre, it drives/or is driven by external entities.

Problems arise when the Business logic is entangled with the external entities.

When viewing the application from inside out we see it communicates with external entities over ports.
Port in this case is meant to evoke thoughts of ports in an operation system.
Any device which adheres to the protocol of the port can be plugged in.

The protocol of the port is given by the *purpose of the conversation* between the two sides.
The protocol takes the form of an API

For each external device there is an adapter that converts the API definition to the signals needed by the device, and vice versa.
- GUI is an adapter which maps movements of a person to the API of the port.
- Storing of data is a port, which can have any number of adapters (SQL DB, flat file)

## Structure

At the centre is the application, inside the hexagon.
Inside the hexagon are just the things that are important for the business problem, i.e. the business logic.
No references to any technology, framework or real world device. So, the app is tech agnostic.

Outside the hexagon we have any real world thing that the app interacts with. These are called the actors, humans, other apps, hw/sw devices.

Actors are arranged around the hexagon.
left/top are drivers, interaction is triggered by the driver.
right/bottom are driven, interaction is triggered by the application.
Two kinds of driven actors:
1. Repository
    store/obtain information
2. Recipient
    Send and forget

On the boundaries of the hexagon are the ports/adapters
They are organized by the reason they are interacting with the app.
Each group of interactions with a given purpose/intention is a port.

### Ports

Ports should be named according to what they are for, not according to any technology.
Use verbs ending in "...ing"
"This port is for ...ing something"

Ex: This driver port is for accepting a service order.
    This driven port is for storing information about a service order

Driver ports are the use case boundary of the application, they are the API of the app.
We can have a port interface with many use cases or just a few.

Single Responsibility Principle - a lot of ports , each one for a use case.
or
Command Bus for a port, with a command handler for each use case.

Driven ports is an interface for functionality needed by the app.
SPI (Service Provider Interface)

### Adapters

Actors interact with hexagon ports through adapters using a specific technology.
There can be many adapters for one port.
Adapters are outside the application.

For each port there should be at least two adapters.
1. for the real driver
2. for testing the behaviour of the port

What an adapter does in the end is convert an interface into another, so we could use the Adapter design pattern to implement it.

Which adapter to use is configured at app startup, this enables mock adapters to be put in for testing.

### Composition Root

This component runs at startup and builds the whole system
* initializes and configures the env: dbs, servers, ...
* for each driven port, it chooses the adapter and creates an instance of it.
* creates an instance of the application injecting the driven adapters
* for each driver port
  * it chooses the adapter and creates an instance
  * runs the driver adapter instance.

### Overall structure

He is recommending different maven modules.
One for the hexagon, one module for each adapter and one module to start the whole project.
The hexagon should only expose ports, so in java you would end up using the Java Module system to only expose the ports.
If the language doesn't support only exposing ports then the hexagon would have to be split into ports and implementation(business logic) modules.

## pros and cons

### pros

Testability improvement
Maintainability improvement
Flexibility
apps immune to tech evolution
delay tech decisions

### cons

Complexity
Build process performance
indirection and mappings
