# microprofile

Eclipse MicroProfile

open source community spec for Cloud Native Java microservices

accelerate the adoption of microservices in the Enterprise java community

different to JCP

CDI is at the core of microprofile

each release brings in more of Java EE
https://projects.eclipse.org/projects/technology.microprofile/

1.0
    CDI
    JSON-P
    JAXRS

1.1
    Config (configure microservice without repacking app, 12 factor app)

1.2
    fault tolerance
    jwt (RBAC for microservice endpoints)
    health check (same as hystrix and failsafe libraries)
    metrics

1.3
    open tracing
    open api

2.0
    JSON-B
    update some versions

**Doesn't include the application server.**

## fault tolerance

@Retry
@Timeout
@Bulkhead (limit the number of concurrent requests)
@CircuitBreaker
    fail on exception
    open => no requests go through
    after a set timeout, half close to let a small number through
    if they're successful close they're circuit
@Fallback
    exception handler???

istio
    An open platform to connect, manage, and secure microservices
    https://istio.io/