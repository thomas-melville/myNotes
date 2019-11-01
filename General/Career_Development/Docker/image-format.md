# image format/structure

https://medium.com/tenable-techblog/a-peek-into-docker-images-b4d6b2362eb
https://github.com/opencontainers/image-spec/blob/master/spec.md

made up of one or more layers and metadata
when you do a docker pull you see all the layers being pulled down

An image is made up of:

1. a file for each layer
  The name corresponds to the hash which is used in the docker pull command
  The name is the sha256 of the contents of the layer
2. a manifest file
  in json format
3. one extra file, config.json
  in json format
  contains metadata about image creation
  filename is it's hash


Docker layers are stored using a Content Addressable Storage Scheme
  the hash of the contents of the layer is how that layer is referred to and stored in the file system.

The manifest file has a layers section
  lists each layer in the image
  with the following information:
    mediaType
    size
    digest (including the algorithm)
  The ordering is important
    docker images use a union file system
      transparently overlay the contents of layers on top of previously layers
        the top layers overwrite lower layers

The manifest also has a config section
  this points to the config file
  same data as layers section

The image digest which we see at the end of a docker pull comes from the sha256 of the manifest.json file
  known as merkle tree
    a tree in which each leaf node is labelled with the hash of a data block
    and every non leaf is labelled with the cryptographic hash of the labels of its child nodes
  An image can be referred to by the hash of its manifest
    and the manifest contains a list of the child dependencies of the image
    If an layer changes, it's digest changes, which changes the digest of the image

You can pull an image by it's hash by appending @sha256... instead of :tag

https://github.com/opencontainers/image-spec/blob/master/spec.md#notational-conventions
