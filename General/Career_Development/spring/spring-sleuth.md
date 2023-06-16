# Spring Sleuth

https://cloud.spring.io/spring-cloud-static/spring-cloud-sleuth/2.1.0.RELEASE/single/spring-cloud-sleuth.html

Implements a distributed tracing solution for Spring cloud.

Spring project used to pass information from service to service.
It's useful for distributed tracing.

The first service that has sleuth starts the trace.

MDC - Mapped Diagnostic Context
  provides a way to enrich logs with information which would not be available in the scope where the logging actually occurs.

Spring Sleuth adds the information to the MDC so it can be automatically added to the log messages.

Brave is used as the tracing library, https://github.com/openzipkin/brave
  Intercepts production requests to gather timing data, correlate and propagate trace contexts.
  This is what interacts with SLF4J
    https://github.com/openzipkin/brave/tree/master/context/slf4j

## Terminology

Span
  the basic unit of work.
  sending an RPC is a span
  identified by a unique 64bit ID and another 64bit ID for the trace the span is a part of.
  Root span starts the trace.

Trace
  A set of spans forming a tree like structure
  A trace can cross microservice boundaries.
  The trace starts when the user initiates the request at the boundary of the system.
  And is carried through all the microservices interacted with using headers.

Span Context
  The state that must get propagated to any child spans across process boundaries.
  Trace and Span IDs are required.
  Baggage is optional

Baggage
  set of k:v pairs stored in the span context
  Baggage travels with trace and span as additional headers.
  prefixed with "baggage-" for header and "baggage_" for message

## Implementation tracing with Brave

brave.Tracer object is used to start/stop traces and spans.
