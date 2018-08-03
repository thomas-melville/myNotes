# RAML

REST API Modelling language

Specify everything about the REST API in YAML and auto-generate the interface with the jax-rs annotations, or whatever library you use
	resource
	queryParameters
	headers
	expected body structure
	what status codes to return
	what to return based on media type
	documentation (which can be auto-generated to API-docs)
	...

2 versions available 0.8 and 1.0

1.0 has more features but tooling support isn't great
springmvc-raml-plugin works grand with 1.0

## language

There's a lot in it.
Just go through the tutorials: https://raml.org/developers/raml-200-tutorial

/uri-part
	http-method(s)
		(for each)
			description
			?queryParameter?
			responses:
				codes
					(for each)
						body
						example

request and response bodies can be json schemas which generate POJOs!!!
these schemas can be extracted and reused

schemas:
	myType: |
	{
		schema ...
	}

...

	type: myType

resource types can create generic collection and collection-items which can be parameterized so they can be reused
specify parameters using this construct: <<...>>
There are special ones, resourcePath and resourcePathName

resourceTypes:
	collection:
		description: blah blah <<...>> blah
	collection-item:

it is possible to change the word by using !pluralize or !singularize

you can specify your own parameters and pass them into the resourceType when you use it

/blah
	type:
		collection-item:
			exampleItem: vbruiawlvnrwe;

in the resource type use it like the reserved parameters

you can extract raml to separate files and include them in the main file using !includes
extracted files are treated as raw strings by RAML
good for examples / schemas/ ...
extract all resourceTypes / traits to a single file

traits

sorting, paging, filtering

creating a trait

traits:
	searchable:
		...

using a trait

/...
	get:
		is: [searchable: {description: "!....", example:"..."}]
