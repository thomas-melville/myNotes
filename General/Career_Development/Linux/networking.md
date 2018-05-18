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