# url en/decoding in java

certain special characters can't be interpreted/transferred correctly so they need to be encoded.
for example spaces

## basic uri syntax

scheme:[user:password@]host[:port][/]path[?query][#fragment]

quick way to see the parts is to pipe a string url in the constructor of a uri

URI uri = new URI("http://localhost:8080/#/contexts/name_of_context?field=datasource")

[] means they are optional

scheme:     what is the transfer means, ex: http ot ftp
host:       the location, ex: www.google.com
query:      the request parameters, ex: http: ?param=x&name=y
fragment:   location to scroll to on the page / path to resource

http://www.baeldung.com/java-url-encoding-decoding states that a common mistake is to encode the complete URI, when typically we only want to encode the query.

URLEncoder.encode(string, encodingScheme)
    only really works for query part
    can't be used for path part of URL

valid characters in query and path parts differ

UriUtils by Spring Framework