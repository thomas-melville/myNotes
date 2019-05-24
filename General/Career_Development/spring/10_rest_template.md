# RestTemplate

https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/web/client/RestTemplate.html
https://www.baeldung.com/rest-template

Synchronous client to perform Http requests, exposing a simple, template method API over underlying HTTP Client libraries.
Default is the JDK Http library, you can use Apache Http, ...
Offers templates for common scenarios by Http method, in addition to the generalized exchange and execute methods that support less frequent cases.

*It will be deprecated in a future version and no new major features. WebClient is the recommended replacement*

## usage

Create an instance, either with new or injected as Bean.
Injected as a bean makes it easier to test, and configure once use everywhere!

```

RestTemplate template = new RestTemplate();

```

Simplest usage is to get plain json

```

ResponseEntity<String> response = template.getForEntity(url, String.class);

```

You can also marshal the json directly into a POJO

```

ResponseEntity<MyPojo> response = template.getForObject(url, MyPojo.class);

```

### Methods specific to HTTP methods

* HEAD
  * headForHeaders(url)
* POST
  * postForLocation
    * returns the URI to the newly created resource.
  * postForEntity
    * returns the resource itself.
  * postForObject
    * returns the resource itself.
* OPTIONS
  * optionsForAllow(url)
* DELETE
  * delete(url)

### generic exchange api

it takes 4 parameters:

* url
* method
* request body
* return type

#### with callback

the request body parameter is wrapped in a RequestCallback

### send form data

The request body should be a LinkedMultiValueMap and the content type: applicatio/x-www-urlencoded

### configuring the RestTemplate

You can pass a ClientHttpRequestFactory into the constructor.
In this class you can configure lots of things:
* timeouts
* ...

You can configure the HttpClient and pass it into the ClientHttpRequestFactory. By default it uses SimpleClientHttpRequestFactory.

You can pass the RestTemplateBuilder class into the Bean.
The builder is injected by Spring and all the autoconfiguration magic happens

Once created you can see a number of objects in it, to overwrite the default behaviour
* ResponseErrorHandler - by default it's DefaultResponseErrorHandler
* ResponseInterceptor
* HttpMessageConvertor
  * There's a number of these: ByteArray, String, Resource, Source, AllEncompassingForm, ...
* RequestCallback
* ResponseExtractor
