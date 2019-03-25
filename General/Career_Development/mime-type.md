# MIME Type

A media type (Multipurpose Internet Mail Extensions) is a standard which indicates the nature and format of data.
The IANA is responsible for all MIME types.

## Identification Structure

The simplest MIME type consists of a type (general category) and a subtype (exact kind of data), each are strings, separated by a /
e.g. application/json
Each type has it's own set of subtypes.
optional parameters can be added.
  An example is charset in text type


## types

There are two classes of types:

* discrete
  * represent a single file / medium
* multipart
  * represents data which comprises of multiple component parts
  * each of which can have it's own MIME type.

### discrete

* application
* audio
* example
* font
* image
* text

### multipart

* message
  * A message that encapsulates other messages
* multipart
  * multiple components that each may have their own MIME type
  * form-data
    * for data produced using the FormData API
  * byteranges
    * for when fetched data is only part of the content

#### multipart/form-data

The document/data/body consists of different parts, delimited by a boundary
the boundary is a string starting with --
each part is it's own entity with it's own HTTP headers
* Content-Disposition
  * State whether the part is to be displayed inline or as an attachment
  * In the multipart form it gives info about the part
  * In HTTP, only the parameter form-data as well as name & filename directives can be used.
* Content-Type
  * Used when the part is a file
