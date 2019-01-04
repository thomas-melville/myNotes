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

## communication between client and server

by default docker runs through a non-networked unix socket.
you can also communicate through HTTPS.
you can enable TLS and point docker to a trusted CA certificate
  the client and server must have certificates in the trusted CA certificate store.

1. on the daemon generate CA private and public keys
2. create server key and certificate signing request
3. sign the public key with our CA
4. Generate the signed certificates

### secure communication between daemon and registry

on the daemon machine under /etc/docker/certs.d/ create a folder for the hostname of the registry
all cert and key files for the registry are placed in this folder/.
