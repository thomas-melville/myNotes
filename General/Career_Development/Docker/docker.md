# docker

docker is a client-server application

docker & dockerd

docker is written in Go and takes advantage of several features of the linux kernel.

## dockerd

a server with a REST API that programs talk to to instruct it what to do
creates and manages docker objects: images, containers, networks, volumes

## docker

a CLI which talks to the docker daemons REST API.

## registry

stores Images
docker hub is a public registry
