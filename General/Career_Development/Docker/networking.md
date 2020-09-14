# networking

## modes

1. bridge
  default mode
  usually for app running in standalone containers
  daemon creates docker0, virtual ethernet bridge
  automatically forwards packets between any other network interfaces that attached to it
  all containers on a host are connected to this internal network
    by creating a pair of peer interfaces
      one to the containers eth0 interface
      other goes in the namespace of the host
    assigns IP address/subnet
2. host
  removes network isolation to the host
  usually for app running in standalone containers
  disables network isolation of a container
  shares network namespace of the host
    may directly be exposed to the public
  inherits ip address of host
3. container
  lets you reuse the network namespace for another container.
  used in kubernetes, multiple containers in a pod
4. none
  no setup from docker
  allows you to set up custom networking

## multi host container networking
