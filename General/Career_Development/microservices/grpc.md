# grpc

communication protocol which is language agnostic
uses protocol buffers to expose and implement stubs
it's like RMI, but not tied to Java

* define a service, methods which can be called remotely with their parameters and return types
* the server implements the service interface and runs a gRPC server to handle client calls.
* the client has a stub of the interface it uses to invoke the server

uses http/2 ontop of TCP