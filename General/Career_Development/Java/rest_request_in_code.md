# sending rest requests in code

Using the Apache HttpClient library

## Post

Create a HttpPost object

```java

    HttpPost post = new HttpPost("host");
```

### multipart/form-data

Example: A request who's body is part json, part file.

Use the **<type>Body** APIs to create the parts of the body with their respective content type
Use the **MultipartEntityBuilder** API to create the body of the request using the parts created.
Make sure to set the boundary in the builder and the content type for the request!
build the entity and give to the HttpPost object

```java

    FileBody fileBody = new FileBody(valuesFilePart.getFile());
    StringBody jsonBody = new StringBody(jsonPart, ContentType.TEXT_PLAIN);
    MultipartEntityBuilder builder = MultipartEntityBuilder.create();
    builder.setBoundary(MULTI_PART_BOUNDARY);
    builder.addPart("json", jsonBody);
    builder.addPart("values", fileBody);
    HttpEntity = builder.build();
```
