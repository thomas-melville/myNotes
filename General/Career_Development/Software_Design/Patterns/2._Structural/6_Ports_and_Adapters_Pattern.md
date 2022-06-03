# Ports and Adapters Pattern

Also known as Hexagonal Architecture

## Intent

Allow an application to equally be driven by users, programs, automated tests or batch scripts.
And to be developed and tested in isolation from its eventual run-time devices and databases.

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
