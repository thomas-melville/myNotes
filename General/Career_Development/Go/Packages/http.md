Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-07-27T13:22:40+01:00

###### http ######
Created Thursday 27 July 2017

http.HandleFunc(<endpoint>, <function>)
	handle the endpoint in question with the function passed.

http.ListenAndServe(<interface>:<port>, ???)
	listen on the interface and port specified
	
the function passed to HandlerFunc
	is of type http.HandlerFunc
	it takes a http.ResponseWriter & http.Request
	ResponseWriter assembles the response, we write to it
	Request is a data structure which represents the request
	
