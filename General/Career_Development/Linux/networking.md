# networking

ipv4 and ipv6

## ipv4 address types

* unicast       associated with a specific host
* network       an address whose host portion is entirely zeros
* broadcast     an address to which each member of a network will listen
* multicast     an address to which appropriately configured nodes will listen

### special address ranges

127.x.x.x           loopback (local system)
0.0.0.0             used by systems which yet don't know their address, DHCP and BOOTP
255.255.255.255     generic broadcast private address, reserved for internal use

other special ranges

10.0.0.0 -> 10.255.255.255
172.16.0.0 -> 172.31.255.255
192.168.0.0 -> 192.168.255.255  local communication within a private network
...

historically IP addresses were defined in classes
Classes A, B, C, ... are used to distinguish the network portion of the address
a netmask is used per class

A   8   24
B   16  16
C   24  8
D   Multicasting
E   not used

network address is achieved by doing a logical && on the IP address and the netmask
This defines a local network which consists of a collection of nodes connected via the same media

## ipv6 address types

* unicast       a packet is delivered to one interface
                **Link-Local**  auto-configured for every interface to have one
                **Global**  dynamically or manually assigned
* multicast     a packet is delivered to multiple interfaces
* anycast       a packet is delivered to the nearest of multiple interfaces
* IPv4-mapped   an IPv4 address mapped to IPv6
* loopback

hostname, just a label to identify a networked device, easier than remembering the ip.
and the ip could change over time.

you can get/set the hostname using the **hostname** command. This does not persist across reboots though
to get it to persist use the **hostnamectl** command
hostname is stored in /etc/hostname

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

**ip** is the preferred command line utility as compared to **ifconfig**
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

nmtui / nmcli for managing connections
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