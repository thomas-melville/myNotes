# nginx

free, open-source, high-perf http server, reverse proxy and IMAP/POP3 proxy server
known for it's high performance, stability, rich feature set, simple config and low resource consumption.

uses event-driven async architecture

## usage

run nginx with a conf file you have edited.

### load balancing options

* round-robin		default
* least-connected	least number of active connections
* ip-hash			use a hash function based on the clients ip address to get sticky sessions

### conf file structure

nginx consists of modules which are controlled by directives specified in the config file

simple directive
	key value pair separated by space and ending in ;
block directive
	key value pair separated by space and ends in {} which contains a list of simple directives

context
	is when a block directive can have block directives inside it.

#### context

events and http contexts reside in the main context
server resides in http
location resides in server

#### server

server {
	listen		the port to listen on
	server_name	the name of the server to send the request to
				_ is a catch all so any request header HOst value will be accepted	
}

#### location

location <path> {
	proxy_set_header <header> <value>
	root <where to send it to>
	proxy_pass <url>
	rewrite <src_path> <dest_path> break
}

path is the path in the request that this location should be matched to
the path is added onto the end of the root which specifies the path to the actual content to serve
nginx selects the location with the longest prefix to match to the request

/api/
/api/reports

if a request with api/reports comes in the it will go to api/reports not api

<url> is the url to pass this request onto

rewrite allows you to change the uri which is passed onto the proxy
you can use a regular expression and matching groups to take part of the src and add it to the dest path

proxy_set_header allows you to set headers to be passed onto the backend
ex: X-Forwarded-For $proxy_add_x_forwarded_for
this adds the clients ip address so to the backend the request came directly from there, not through the nginx

#### upstream

optionally starts with the load balancing mechanism, by default it's round_robin
a block directive which lists a number of ips addresses of multiple backends for load balancing