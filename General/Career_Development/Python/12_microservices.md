# microservices written in python

socket module from python

## flask

for REST based microservice

https://flask-restful.readthedocs.io/en/latest/quickstart.html#full-example

flask uses decorators for defining routes and the functions which react to them!!!
A function can have multiple routes

@app.route

You need to configure WSGI if you want to use flask in production

By default Flask looks for html templates in a folder called templates

@app.errorhandler(404)
  what to return for a specific 4/500 http response code

## django

https://leanpub.com/tangowithdjango19/

for full blown websites
