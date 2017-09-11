Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-01-13T14:52:25+00:00

###### IPv6 ######
Created Wednesday 13 January 2016

The principle feature of ipv6 is the restoration of the e2e model on the back of which the internet was built in the first place.
	e2e connectivity, security, QoS, node reachability, remote access for maintenance, n/w mgmt purposes.
	Which was broken in IPv4 by NAT, Network Address Translation, and private addresses
		necessary evil to meet the demand for new connections before ipv6 was ready and deployed
Other features (tightly) redesigned:
	mobility, multicast, auto-config
Add n/w value to services like RF-Id and sensors.
	IoT, small grids, cloud computing, smart cities, 4G/LTE services
	
**Main technical advantage:** deal with the problem of IPv4 address exhaustion, 192.168.0.90 to FFFF:FFFF:FFFF:FFFF:FFFF:FFFF....(32 bit to 128 bit)
Other ++++
	hierarchial address allocation methods
	elimination of NAT, move to flat address space
		with hierarchical address allocation methods to facilitate route aggergation across the internet
	multicast is expanded and simplified
	improved security
4 & 6 are not interoperable so need [[https://en.wikipedia.org/wiki/IPv6_transition_mechanism|transition mechanisms]]
	6 creates a parallel independent network

8 groups of 4 hex numbers separated by :
	there are abbreviation methods
		don't need leading 0's
		0000 can be omitted in replace of ::
	first 64 bits - subnet prefix
	second 64 - interface identifier

ip command replaces ifconfig

All IPv6 hosts require a link-local address
	prefix: FE80::/10
	right hand side is derived from MAC address
	and 0's filled in between
	
Global addressing
	prefix supplied by router advertisements on the network

Scope can be link, site or global, ...
	link is only local
