# docker-compose

a tool for defining and running multi-container docker applications.
uses yaml

3 step process:
	1. define service env in Dockerfile
	2. define services that make up your app in the docker-compose.yml file
	3. run the docker-compose.yml file
	
## Syntax:

version:
services:
	<container_name>:
		build:
		image:
		ports:
		volumes:
		environment:
		network:
		depends_on:
		container_name:
		...
networks:

### version

the version of docker-compose to use, different versions have slightly different syntax

### services

the list of services to spin up as containers 

#### build

specify the Dockerfile to build

#### image

specify the image to use

#### depends_on

This container depends on another container being up.
Note: up does not mean ready, your application will need to make sure the service it depends on is ready.

### networks

the networks to use to link the containers together.

## CLI

you can specify multiple compose files using the -f flag
settings are hierarchial, files later in the command override information in earlier files

2 things in a docker-compose file
	containers
	networks

networks are created and containers are added to the networks.
containers are given hostnames which match the name of the container in the docker-compose file