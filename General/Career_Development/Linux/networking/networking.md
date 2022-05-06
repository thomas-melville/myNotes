# networking

## osi model

Open Systems Interconnection model, standardizes language used to describe networking protocols.
Each layer only communicates with the layer directly above/below it.
Not all layers are used all the time.

7 layers

Physical
  Light, electric puleses, ...
Data link
  MAC
Network
  IP, ICMP, ...
Transport
  TCP, UDP, SCTP
Session
  reliable socket, ..
Presentation
  Data conversion, SSH
Application
  HTTP, FTP, DNS, ...

### Physical

Signals are converted into information the system can use.
* electric pulses over copper cable
* laser pulses over fibre optic cable
* freq modulation over radio waves

Lowest layer

USB, Wifi, Bluetooth

A unit of data is a frame

### Data link

Accepts data from the hw in layer one.
prepends an address to all inbound packets as it accepts them.
Address number is 6 bytes, 3 bytes for manufacturer and 3 random bytes assigned to adapter.
Known as the MAC address, Media Access Controll address.

Deals with transferring data between network nodes using the MAC address.
In this layer IP address and MAC address associatd using the ARP Protocol. Address Resolution Protocol.

A broadcast is initiated to find the MAC address of the IP address that is required.

#### Bridge

At this layer a bridge can be created, accepts packets on one side of the bridge and passes them through to the other side. also bi-directional.
Have been in the kernel for a long time, can be seen when you are creating containers or vms on a machine.
Often DHCP services are associated with a bridge.
A number of tools available:
* iproute2
* bridge-utils
* netctl
* systemd-networkd
* NetworkManager
* Applications that create virtual environments

### Network Layer

Routing and forwarding packets to the next hop.
This is the backbone of the internet.

Some protocols at this layer:
* IPv4/6
* ICMP, Internet Control Message Protocol
* IGRP, Interior Gateway Routing Protocol

### Transport Layer

Responsible for end to end communication protocols. Data is properly multiplexed by defining source and dest port numbers.
Also deals with reliability by adding checksums, repeat requests and avoiding congestion

TCP, UDP and SCTP are common protocols at this layer

Servers use a fixed port number and clients use random port numbers.

Three different port types:
1. Well known ports
    0-1023
    Assigned by IANA
    ex: 22-ssh, 80
    require super user priviliges to be bound
2. Registered ports
    1024-49151
    assigned by IANA
    ex: 1194 TCP/UDP, OpenVPN 1293
    don't require super user priviliges
3. Dynamic/Ephemeral ports
    49152-65535
    used as source ports for client side of TCP/UDP

### Session Layer

Layer 5 is used for establishing, managing, synchronizing and termination of application connections between local and remote machines.
If a connection is lost or disrupted this layer may try to recover the connection.
If a connection is not used for a long time this layer may close and reopen it.
Two types of sessions:
* connection-mode service
* connectionless-mode service
Session options
* Simplex or duplex communication
* Transport layer reliability
* checkpoints on data units
Session services may be involved in:
* authentication to Transport Layer
* Setup and encryption initialization
* support for streaming media
* support for smtp, http and https protocols
* SOCKS proxy

RPC-type protocols depend on this layer:
* NetBIOS
* RPC
* PPTP

### Presentation layer

Commonly rolled up into a different layer, 5 or 7.
Ex: HTTP Protocol has methods for converting character encoding.

### Application Layer

Common protocols at this layer:
* HTTP
* SMTP
* DNS
* FTP
* DHCP

## network devices

network devices don't have associated device files, they are referred to by their names.
and a number

* eth, enp  ethernet devices
* wlan      wireless devices
* br        bridge interfaces
* vmnet     for virtual devices communicating with virtual clients
* docker    interface for communicating with docker containers
* lo        loopback localhost interface

linking interfaces to physical devices has been problematic
historically, the MAC address of the device was manual tied to an interface

To tackle this there is now the Predictable Network Interface Device Names (PNIDN)
It is stongly correlated with udev and integration with systemd

There are 5 types of names that devices can be given

* Incorporating firmware or BIOS provided index numbers for on-board devices.
* * ex: eno1
* Incorporating firmware or BIOS provided PCI Express hotplug slot number index:
* * ex: ens1
* Incorporaring physical and/or geographical location of the hw connection:
* * ex: enp2s0
* * this is the format of the LAN connection on my laptop!
* * the relevance of the numbers can be seen in the **lspci** output, the first triplet is the bus, device/slot and function
* Incorporating the MAC address:
* * ex: enx7837d1...
* Using the old classic method:
* * ex: eth0

## ip

**ip** is the preferred command line utility as compared to **ifconfig**.
part of the iproute2 package.
more versatile and efficient because it uses **netlink** instead of **ioctl**
can be used to configure, control, query devices and i/f parameters, manipulate routing , policy-based routing and tunnelling

## ip command objects

* address   ipv4, ipv6 devices
* link      network devices
* maddress  multicast address
* monitor   watch for netlink messages
* route     routing table entry
* rule      rule in the routing policy database
* tunnel    tunnel over ip

For each command objects there are a number of flags

* show
* add
* set
* ...

## ifconfig

sys admin utility long found in unix based systems
configure, control and query network interfaces
superseded by **ip**

## NIC config files

ip and ifconfig chages aren't persisted.
A number of linux distro dep config files exist

## Network Manager

is a daemon with a D-bus for communication to applications.
nmtui / nmcli / mngui for managing connections
this is old school
it's a lot more dynamic today

## routing

process of selecting paths in a n/w along which to send traffic
routing table is a list of routes to other networks managed by the sys

### commands

route

ip route

default route is the way packets are sent when there is no other match in the routing table
can be obtained dynamically by using DHCP

static routes are used to control packet flow when there is more than one router or route.
defined for each interface and can be persisted or non-persisted

## Name resolution

translating ip address to hostnames

2 ways

* static name resolution, /etc/hosts
* Dynamic Name Resolution, DNS

on the command line there are 3 ways to get the ip from a hostname

* dig
* host
* nslookup

/etc/hosts is checked before DNS
should only contain mappings for small internal networks

DNS is dynamic and consists of a network of servers
which clients use to look up names
service is distributed
each one has it's own zone of authority

## Network Diagnostic Utilities

### ping

send packets to host to see if it's contactable

### traceroute

display a network path to a destination
by default it sends 3 packets, use -q 1 to only send one

It uses the TTL header and an ICMP message to get the hops.
When TTL is 0 the packet is discarded and an ICMP message it returned to the sender, with the ip address of the current router.

* traceroute starts with a TTL of 1.
* router receives packet,
* decrements tutorials
* sees the value is 0
* discards the packet and returns an ICMP message
*We now have the first hop*
* It repeats this process incrementing the starting TTL every time until it reaches the destination

### mtr

combines the func of ping and traceroute and creates a continuously updated display, like top

### dig

useful for testing DNS functionality

## Firewalls

a firewall is a network security system that monitors and controls all nw traffic
applies rules on both incoming and outgoing network connections and packets
builds flexible barriers depending on level of trust given to connection.
Can be hw or sw based
found in both routers and individual computers / network nodes.
many firewalls having routing capabilities

3 generations of firewalls

* packet filtering
* * contents of each packet inspected
* stateful filters
* * also examines the connection state of the packet
* * new / existing / none
* Application Layer Firewalls
* * are aware of application and protocol

### Packet Filtering

intercepts packets at one or more stages/layers in the network transmission
firewall establishes a set of rules by which each packet may be:

* accepted/rejected
* mangled
* redirected
* inspected for security reasons
* ...

### configuring a firewall

command line:

* iptables
* firewall-cmd
* ufw

GUIs:

* system-config-firewall
* firewall-config
* gufw
* yast

#### firewalld

Dynamic Firewall Manager
utilizes network/firewall zones which have defined levels of trust
supports runtime and permanent changes to config
configuration kept in **/etc/firewalld**

command line version is firewall-cmd, 200 lines of help but args read like sentences

##### zones

each has a defined level of trust and certain known behaviour for traffic
each interface belongs to a particular zone
a zone can be bound to network interface and network addresses, use netmask
any packet not associated with a zone goes to public

* drop
* * all incoming packets dropped with no reply
* * all outgoing connections permitted.
* block
* * all incoming nw connections are rejected
* * only permitted connections are those from within the system
* public
* * do not trust any computers on the network
* * only certain consciously selected incoming connections are permitted.
* external
* * used when masquerading is being used
* * usually in routers
* * trust levels are the same as public
* dmz
* * used when access to some, but not all, services are to be allowed to public
* * only particular incoming connections allowed
* work
* * trust, but not completely, connected nodes to be not harmful
* * only certain incoming connections allowed.
* home
* * mostly trust the other network nodes
* * only certain incoming connections allowed.
* internal
* * similar to work
* trusted
* * all network connections are trusted.

You can get more fine grained and only allow particular services and ports in a zone

firewall-cmd --zone=xxx --list-services / --add-service
firewall-cmd --get-zones
...

## NTP

Network Time Protocol. Keep time consistent across the network.
NTP time sources are divided into 4 strata.

strata 0 - clock is a special purpose time device (atomic clock, GPC radio, etc)
strata 1 - a NTP server which is connected directly to a strata 0 source (over serial or the like)
strata 2 - a NTP server which references a strata 1 server using NTP
strata 3 - a NTP server which references a strata 2 server using NTP

NTP may function as a
* client - acquire time from a peer
* server - provides time to a client
* peer - syncs time between other peers.

## HTTP

### Caching

help reduce perceived lag, network utilization and may improve performance of web apps.
Can also be used as a filtering proxy, restricting access to certain sites or resources.
Two flavours.

* Forward cache
  * speed up HTTP access in a network.
  * When multiple browsers hit the same cache for the same content, it will be returned instead of going all the way to the original source.
  * squid, tinyproxy, apache.
* Reverse cache
  * speed up perceived lag from a HTTP app server to a client.
  * squid, nginx, tinyproxy, apache.

#### Proxy SSL

When a browser fetches a https:// resource, one of the following happens:
* a CONNECT method request is made to the proxy server, and traffic is transparently forwarded to the destination.
* browser bypasses ther proxy and goes directly to the destination.
* Sslbump peek and splice - makes bumping decisions after the origin server is known.

#### Cache hierarchy

further extends the cache idea. Groups of cache servers working in concert can increase cache efficiency, route the traffic to the best link and support a higher number of clients.
Two types:
* Peer to Peer
  * ask all or some of their peers if they already have the cached content
  * if not the cache server requests the content itself.
* Parent/Child
  * asks a parent if it has the content
  * parent then fetches the content on behalf of the child.

These two types can be intermixed.

## DMZ

Demilitarized Zone.
Special purpose network housing business critical servers which need access to a large untrusted network.
