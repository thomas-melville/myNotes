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

### registry structure

At the root of the registry are two folders

* blobs
* repositories

#### repositories

This folder contains a sub folder for each repository.
A repository stores multiple versions of a single image.

Each repository will have three sub folders.

* \_layers
  * contains a folder which named the checksum of each layer
  * each folder contains one file, which links to a specific file in the blobs folder.
* \_manifests
  * contains two folders
    * revisions
      * contains a folder which is named the checksum of manifest file containing a file called link
      * this link file points to a file in the blobs folder
    * tags
      * contains a folder for each tag of the image
      * each folder contains two sub folders
        * current
          * which contains a link file back to the manifest file in the blobs folder.
        * index
          * which contains a link file back to the manifest file in the blobx folder.
* \_uploads
  * a temporary folder used by docker during a push of an image.
  * Registry has a background process to clean it up.
  * When you initiate an upload a uuid folder is created which is used during the upload.
  * Once the upload is complete the data is moved into the blobs directory under a content addressable scheme and the uuid folder is deleted.

#### blobs

contains the actual image data

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
