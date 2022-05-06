# ip protocol

Two functions:

1. Addressing
    exmaines the address of the incoming packet and decides if the datagram is for the local system or another system
    if local, headers removed and datagram is passed up to the next layer.
    if not, then datagram is passed to next system in direction of final destination
2. Fragmentation
    split and reassemble packets if the path to the next system uses a smaller transmission unit size.

## ip address parts

network and host

a netmask is used to distinguish between the two parts

A subnet mask can be used to break the host part down into further sub networks

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

Not the best use of addresses so CIDR (Classless Inter-Domain Routing). Uses a number bit mask instead of a class bitmask.

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
