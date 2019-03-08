# spring cloud contract

https://cloud.spring.io/spring-cloud-contract/single/spring-cloud-contract.html

umbrella project for implementing Consumer Driven Contract testing.
One project in it at the minute:
* Spring Cloud Contract Verifier

Enables CDC development for JVM based apps.
It's shipped with a groovy/yaml DSL

Contract definitions produce the following resources:

* json stub definitions to be used by wiremock
  * for doing integration testing on the client side
* acceptance tests
  * junit or spock
  * verify that the server side impl of the API is compliant with the contract

## terms

producer - of the API for messaging
consumer - of the API for messaging
server - of the API for HTTP
client - of the API for HTTP

## steps

1. add the maven dependencies
2. create a base class
3. create the contract
4. add the plugin to generate the tests
    * configure base class here

## base class

contracts can share the same base class or each can have it's own base class or a mix
the last two folders in the package structure determines the base class it will extend.
and the class will end in base
You have to write the base classes to mock what we need to

## contracts

the parameters in the request can be specific values or a match regex.

## DSL

what's with consumer() & producer() methods?
  there are two users of the contract the product and consumer.
  The value expected/returned is different based on whether it's the consumer, going into the stub, producer, an assertion in the unit test.

description
  one liner of the purpose of the contract
  and the specifics in a given, when, then format
  it's just arbitrary text though

name
  you can give the contract a name
  this is used to generate the test method name
  Otherwise the test method name comes from the groovy file name

file
  you can pass in the contents of a file to a contract
  request/response body seem the most obvious
  the path is relative to the current contract

fileAsBytes
  pass the contents of the file as bytes

request
  a closure which allows you to specify the request
  mandatory are:
    method
    urlPath
  you can specify absolute path using url, but the relative urlPath is the recommended way

response
  a closure which allows you to specify the response
  must contain a http status code
  may contain headers, cookies and body

urlPath
  takes a string
  and is a closure
  which can take the following:
    queryParameters closure
    headers closure
    cookies closure
    body closure
      only json is supported at the minute
    multipart elements

queryParameters
  is a closure which takes a list of parameters:
    parameter 'key' : 'value'
    the value here can be many things
      a simple literal value
      an assertion to a value, matching, containing, ...
      a call to value which allows for values specific to the consumer & producer

headers
  a closure to specify headers for the request/response
  individual headers can be:
    header 'key' : 'value'
    dsl methods, contentType, ...

multipart
  a method which can take multiple arguments
    formParameter
    booleanParameter
    file
      takes an argument which is a method called named
      has 3 arguments
        name
        content, can be a string or point to a file
        contentType

execute
  allows you to execute a groovy method and use the value
  you can also define your own methods in the base class and call them in this

optional
  allows the consumer side of the request or the producer side of the response to be optional
  i.e. no need to pass a value

fromRequest
  get access to a lot of information from the request
  so you can use it in the response

### dynamic properties

timestamps, ids, ...
two ways in the groovy DSL
  directly in the body
  body matchers

value(consumer(), producer())
you can use regular expressions to allow the consumer to enter a range of values

you can use bodyMatchers for two reasons
  define dynamic values that should end up in the stub
  verify the result of your test.

lots of ways
  by...()
