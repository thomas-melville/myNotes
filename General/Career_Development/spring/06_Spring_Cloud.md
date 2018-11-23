# Spring Cloud Fundamentals

Spring cloud helps you build cloud native applications.
It is a number of projects which fall under the same umbrella.
Built on top of Spring Boot

* Spring Cloud Config
* Spring Cloud Cluster
* Spring Cloud Stream
* Spring Cloud Netflix
* Spring Cloud Sleuth
* Spring Cloud Bus
* Spring Cloud Consul
* ...

Focus of pluralsight course

* Service Discovery
* Distributed Configuration
* Intelligent Routing
* Client side load balancing
* Circuit breaker

## Service Discovery

### What is service discovery

In the new architecture, decomposition into multiple services which can scale independently.
How does one service know where all the other services are, when they can come and go on a massive scale.

Provides the following functionality:

* A way for a service to register itself
* A way for a service to deregister itself
* A way for clients to find other services
* A way to check the health of a service

There are a number of projects you can use to discover services:

* Spring Cloud Consul
* Spring Cloud ZooKeeper
* Spring Cloud Netflix

Spring Cloud Netflix = Netflix OSS + Spring + Spring Boot, i.e. a collection of projects again

* Eureka Server
* Eureka Client
* ...

There are 3 components involved in Service Discovery

* Discovery Server
* Service
* Client

The process:

* service registers location
* * service identifier
* * location, ip or dns
* * port
* client looks up service location
* discovery service sends back location and port
* client requests service at location
* service responds

### Discovery Server

Actively managed registry of service locations.
The source of truth
run more than one instance
Spring Cloud Eureka Server
When it starts up it tries to register itself with a peer server. To turn this off use the property **eureka.client.register-with-eureka**
By default the discovery server runs on port 8761

* add dependency in dependencyManagement
* * org.springframework.cloud:spring-cloud-dependencies
* * type = pom, scope = import
* add dependency in dependencies
* * org.springframework.cloud:spring-cloud-starter-eureka-server:...
* add property:
* * spring.application.name=discovery-server
* In main app class add one anotation
* * @EnableEurekaServer

Now you have a running instance of a discovery server

### Application Service

user of the discovery client

* add dependency in dependencyManagement
* * org.springframework.cloud:spring-cloud-dependencies
* * type = pom, scope = import
* add dependency in dependencies
* * org.springframework.cloud:spring-cloud-starter-eureka:...
* add two properties:
* * spring.application.name=my-service
* * eureka.client.service-url.defaultZone=http://localhost:8761/eureka
* * * the zone thing is because of AWS
* In main app class add one annotation
* * @EnableDiscoveryClient

### Application Client

depends on other services
also a user of the discovery server
A service and be a service and a client of the same discovery server

* same setup as application service
* add one more property:
* * eureka.client.register-with-eureka=false
* Inject discovery client
* * can be Eureka specific or Discovery server agnostic
* * @Inject EurekaClient client / @Inject DiscoveryClient client
* Get service instance info
* * eurekaClient.getNextServerFromEureka(<service_name>, <is_secure_request>)
* * * RoundRobin
* * discoveryClient.getInstances(<service_name>)
* * * you get the full list.
* * **by default apps use the spring.application.name as their virtual hostname when registering**

```java

//...

    @Inject / @Autowired
    DiscoveryClient client;

    //...

        InstanceInfo info = client.getInstances("service");

        String baseUrl = info.get(0).getUri().toString()

```

### Eureka

Provides a dashboard which shows useful data about the available services

eureka.dashboard.enabled=true

#### configuration

eureka.server.*

configuration the discovery server

eureka.client.*

how the client interacts with the server

eureka.instance.*

how the instance registers itself with the server, information it provides to the server

#### health & availability

regularly checks the status of the services
heartbeat every 30 seconds
default behaviour, you can configure it to it an endpoint, for example /health in spring boot apps
if no response in 90 seconds service is removed

discovery server sends back a copy of the registry to a client
so that the clients can operate if the server goes down
clients then fetch deltas to update their registry

#### AWS

Eureka heavily supports AWS

It can detect if it's running on an AWS instance
if so it calls out and gets some info
tries to bind to next available elastic ip
multi-zone aware
multi-region aware

some extra config:

* Create a config class
* Then a bean, EurekaInstanceConfigBean
* add some amazon stuff to it, AmazonInfo

lots of specific AWS configuration properties based on regions and zones

## Configuration Distribution

Achieved using a Configuration server.
A dedicated, dynamic, centralized k/v store, may be distributed
authoritative source
auditing
versioning
cryptography support

There are a number of projects you can use to manage configuration:

* Spring Cloud Consul
* Spring Cloud Zookeeper
* Spring Cloud Config

Consul & Zookeeper are more than just config servers, they are for service discovery also.
Spring Cloud Config is a dedicated config server and client.

Client:

* embedded in application
* Fits seamlessly into Spring Environment abstraction
* responsible for bootstrapping & fetching config
* * bootstrapping has to be done very early, even before the ApplicationContext is created.
* * two ways:
* * * config first
* * * * specify location of config server
* * * discovery first
* * * * discover location of config server

Server:

* standalone application
* Fits seamlessly into the Spring PropertySource abstraction
* HTTP REST Access to config files
* Write files, name them accordingly
* server is only read, no write
* can handle a number of formats
* * default JSON
* * properties
* * yaml
* Backend Stores
* * Git
* * SVN
* * Filesystem
* Configuration scopes
* * global across all apps
* * app specific
* * profile specific

get up and running

* add dependency in dependencyManagement
* * org.springframework.cloud:spring-cloud-dependencies
* * type = pom, scope = import
* add dependency in dependencies
* * org.springframework.cloud:spring-cloud-config-server
* create a folder to store configurations
* This can optionally contain a application properties file which contains global properties
* add app & profile specific files
* run git init
* git add, commit
* setup remote

the config server can then push of updates when changes are committed to the repo

* server.port=8888 convention
* spring.cloud.config.server.git.uri=<...>
* Add an annotation to the main class:
* * @EnableConfigServer

(Add the eureka client to make the config server discoverable by other services)

using the server

* add dependency in dependencyManagement
* * org.springframework.cloud:spring-cloud-dependencies
* * type = pom, scope = import
* add dependency in dependencies
* * org.springframework.cloud:spring-cloud-config-client
* create bootstrap.yml|properties:
* * spring.application.name
* config first:
* * spring.cloud.config.uri
* discovery first
* * spring.cloud.config.discovery.enabled

### REST Endpoints

parameters:

* application
* * maps to spring.application.name
* profile
* label
* * a feature for grouping config files
* * in git backend this is the branch

GET /{application}/{profile}[/{label}]
GET /{application}-{profile}.[yml | properties]
GET /{label}/{application}-{profile}.[yml | properties]

### updating config at runtime

refresh config without restarting the app

* refresh @ConfigProperties
* Update logging levels
* Any @Bean / @Value annotated with @RefreshScope

**If a bean takes an instance of a ConfigurationProperties, the bean instance won't be refreshed without @RefreshScope.**

You can put RefreshScope on the whole app, specific classes or specific fields

1. update the config and push the change to the repository which the config server points at

Multiple ways to refresh the config

1. /actuator
* * call the refresh endpoint from the actuator
* * need to do this for every instance
* * I've done this programmatically using a cron
2. /bus/actuator
* * spring-cloud-bus
* * message bus which all apps subscribe to
* * send event and all apps will refresh
3. VCS + /monitor
* * spring-cloud-config-monitor + spring-cloud-bus
* * link remote git repo to /monitor endpoint, repo posts to endpoint when there's a change

### encrypting and decrypting configuration

supports encrypted configuration at rest and in flight, and can use symmetric or asymmetric keys

symmetric keys are easier to use but less secure.
asymmetric keys, private & public keys

It /encrypt and /decrypt endpoints, but neither are secure

The configuration can be decrypted at 2 points:

* by the config server before it sends it
* locally at the client on response

1. choose your key type
2. configure config server
* * symmetric
* * * add key to config server
* * asymmetric
* * * add pem as text to config server
* * * or
* * * add keys to java key store
* * * point config server to key store

## Intelligent Routing

How to simplify the interface when there are multiple service instances which are on different ips, ports and have different api paths

Implemented by a Gateway Service or API Gateway

"Single point of entry for all clients"

* Dynamic routing & delivery
* security
* filtering
* Auditing
* logging
* request enhancement
* load balancer
* different apis for different clients

spring-cloud + netflix Zuul

gateway service that provides dynamic routing, monitoring, resiliency, security & more

* add dependency in dependencyManagement
* add dependency in dependencies
* in main app class as @EnableZuulProxy
* * This app then becomes the gateway service
* it can use Service Discovery the usual way
* configure request routing
* * default behaviour when service discovery is enabled is to route by service name
* * by default the prefix is stripped from the request
* * using properties you can configure more precise routing

client side load balancer project from Netflix is called ribbon

### filters

different types:

* pre
* * before the request
* route
* * direct the request
* post
* * afteer the request
* error

extend and implement the abstract methods in the ZuulFilter class

* run, what to actually do. the course just adds headers to it
* shouldFilter, a boolean whether to run the filter or not
* filterType, a string representing pre, route, post or error
* filterOrder

RequestContext object holds state information about the request

Add your implementation to the App Context as a Bean

## Load balancing

two types

* Server side
* * service in front of instance which balances the load
* * typically hardware based, but can be software
* * additional hop
* * support various algorithms
* * happens outside the request process
* * centralized / distributed
* Client side
* * client is aware of instances and balances the load
* * software based
* * support various algorithms
* * happens inside the request process
* * distributed

Client side is a natural fit for cloud native apps

Netflix Ribbon: "is a inter process communication (rpc) library with built in software load balancers"

Spring cloud adds full integration with Ribbon in RestTemplate class
RestTemplate will know how to balance requests across multiple instances of a service
you can customize the algorithm
round robin by default

* add dependency mgmt on spring cloud
* add dependency on spring cloud ribbon
* New annotations
* * @LoadBalanced
* * * marks a RestTemplate to support load balancing
* * * Do this on a bean in a configuration class
* * @RibbonClient
* * * used for custom configuration and when Service Discovery is absent

This should be combined with Service Discovery

restTemplate.get...("http://service-name/...")

RestTemplate will resolve "service-name" to an ip and port from the service discovery and load balance it

Without Service Discovery you need to define the Ribbon Client

* Add annotation on main class, @RibbonClient(name = "some-service")
* Add some properties with the prefix "some-service"
* * <...>.ribbon.eureka.enabled=false
* * <...>.ribbon.listOfServers=... , ...

you then use the name of the ribbonclient in ther url for the rest template

you can add more configuration to the RibbonClient annotation:
    configuration=MyConfiguration.class
The class should be outside the component scan base
There are a number of beans you can specify in this configuration class

* IClientConfig
* IRule
* * control load balancing strategy
* * You can create your own but there are some existing impl
* * RoundRobinRule, ResponseTimeWeightedRul, RandomRule, ZoneAvoidanceRule
* IPing
* * used to check the liveliness of a service
* * You can create your own but there are some existing impl
* * DummyPing, PingUrl, NIWSDiscoveryPing
* ServerList<Server>
* ServerListFilter<Server>
* IloadBalancer

## Self-healing services

Netflix Hystrix & Turbine

In a distributed system failure is inevitable

hardware, networks, software
more of everything
process communication, is now over a network!

cascading failures
domino effect
a failure in one causes all others to fail

fault tolerance and gracefully degrade or fail fast
resource overloading

Hystrix is an implementation of the circuit breaker pattern
wraps calls and watches for failures

10 second rolling window
20 request volume
>= 50% error rate

waits & tries a single request every 10 seconds
failed requests can execute fallback methods
all requests run in thread pools

* add dependencyMgmt spring-cloud-dependencies
* add dependencies spring-cloud-starter-hystrix
* annotate method with @HystrixCommand
* * add a field, fallbackMethod, and specify a method name

be careful with hystrix timeouts
ensure timeouts encompass caller timeouts & any retries, default 1000milliseconds

### Hystrix Dashboard

circuit state
error rate
traffic volume
successful, rejected requests
timeouts
...

Nice looking dashboard

* add dependencyMgmt spring-cloud-dependencies
* add dependencies spring-cloud-starter-hystrix-dashboard
* Add annotation @EnableHystrixDashboard
* start it up and enter what app you want to monitor, http://...:.../hystrix.stream

Turbine, aggregate many hystrix streams into one

* add dependencyMgmt spring-cloud-dependencies
* add dependencies spring-cloud-starter-turbine
* @enableTurbine
* application.properties
* * turbine.app-config= list of service ids
* * turbine.cluster-name-expression=...(name cluster)