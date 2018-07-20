# networking commands

## hostname

displays information about the hostname

d domain name
f fully qualified name
i ip address

## ping

send packets to a host to verify:

1. it is reachable
2. the speed of the connection

### flags

c<x> transmit x packets then stop

## ifconfig

display the network config of the machine.
All the interfaces it has to the network.

## netstat

finding connections to and from the host
find out what application is listening on a port

netstat -nap | grep 80
    find out the PID of the process listening on port 80

lsof can achieve the same thing

lsof -i :80

### arguments

* n: show numeric addresses, instead of trying to determine symbolic host, port or user names.
* t: tcp connections
* l: show only listening sockets (these are omitted by default) 
* p: show the pid and name of process to which each socket belongs

## nslookup

find the hostname of an ip and vice versa

## traceroute

view number of hops and response times to a remote system.

## finger

view user information for users on a system

## telnet

connect over the telnet protocol, use to see if a host is alive