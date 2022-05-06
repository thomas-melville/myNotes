# remote access

## telnet

one of the earlier protocols developed back in 1969.
No security, data is sent over the wire in plain text.

## rsh

written for BSD in 1983.
similar to telnet.
Is also unsecure.

## ssh

Secure Shell protocol.
developed to overcome the security concerns of telnet & rsh.
Most widely used version of the protocol is OpenSSH.
Built upon the concepts of symmetric and asymmetric encryption.
Multiple layers of the OpenSSH architecture handle different parts of the connection.
Other protocols can be tunneled over SSH: X11, VNC.

openSSH key-based Authentication. ssh-agent can cache decryped private keys. ssh-copy-id can copy your public key to a remote host.

### Transport Layer

Deals with the initial key exchange and setting up symmetric-key session.

### User Authentication Layer

Deals with Authenticating and authorizing user accounts.

### Connection Layer

Deals with the communication once the session is setup.

### Server

Configuration is in /etc/ssh/sshd_config

### Client

Configuration is in /etc/ssh/ssh_config (Note no, d after ssh)

You can setup shortcuts to servers you usually access, $HOME/.ssh/config

### openSSH Tunnel

#### Local Tunnel

-L - indicates which port to be opened on the local host and the final destination. The connection to the final destination is going to be made by another machine.

ssh -L 4242:host-c:220 host-b

#### Remote Tunnel

-R - requests another machine to open a listening port to which any connection will be transferred to the destination.

ssh -L 4242:host-c:220 host-b

#### Dynamic Port Forwarding

-B - create a socket on the local machine, which acts as a SOCKS proxy server. When a client connects to the port it is forwarded to the remote machine, which is then forwarded to the dynamic port on the destination machine

### parallel ssh commands

Often, it is required to execute the same command on multiple machines. pssh package!

* pssh - parallel ssh
* pnuke - parallel process kill
* prsync - parallel copy program using rsync
* pscp - parallel copy using scp
* pslurp - parallel copy from hosts.

The commands use the existing ssh configuration. Best to configure aliases before using pssh.
If there is a password, or fingerprint prompt the pssh command will fail.
Convenient to put ip addresses in a file and pass that into pssh commands.

## VNC

Virtual Network Computing (VNC) Server allows for cross platform graphical remote access.
Several implementations:
* tigervnc
* Xvnc
* vncserver
    opens ports from 5901 up
* vncpasswd
* vncvonfig
* vncclient
    display-based protocol, making it cross platform
    heavy protocol, as pixel updates have to be sent over the wire.

On it's own VNC is not secure. However, it can be tunneled through SSH or VPN connections.

## X Window System

Simple network transparent GUI system provides basic GUI primitives.
User interfaces as provided by toolkits and add-ons.

Uses the X11 protocol.

Depends on the display variable.
Authentication is done using keys. Keys are managed by xhost or xauth.

To secure the connection can be tunneled with VPN or SSH, OpenSSH supports X11 tunneling.
use the -X argument to tunnel an X11 session over ssh

## DNS

Before DNS, there was ARPANET.
The original solution to a name service was a flat text file, HOSTS.txt. Hosted on a single machine, and when you wanted a copy you ftp'd it from the central machine.
Did not scale well!

A descendant of the HOSTS.txt file is the hosts file on Linux, in /etc/hosts

<IP Address> <Hostname> [Hostname or Alias]

DNS is a distributed, hierarchial database for converting DNS names into IP Addresses.
The k-v store can be used for more than just IP address information.

DNS protocol runs in two modes:
1. Recursive with caching
2. Authoritative

Authoritative mode is used in the root DNS and recursive is used in all sub servers in the hierarchy.

Database is broken into tree nodes called Domains. Domains are managed as part of a zone.
Zones are the area of the namespace managed by the Authoritative server. DNS delegation is done on Zone boundaries.

Example:

DNS request for host1.foo.example.com

client makes a request to DNS server it has configured to talk to.
caching nameserver makes a query for host1.foo.example.com to the root **.** zone nameserver.
the **.** root zone nameserver points to the nameserver for the **com.** zone.
caching nameserver makes a query for host1.foo.example.com to the nameserver for the **com.** zone.
The **com.** zone nameserver sends a response that points to the nameserver for **example.com** zone.
caching nameserver makes a query for host1.foo.example.com to the nameserver for the example.com. zone.
The **example.com.** zone nameserver sends a response that points to the nameserver for **foo.example.com** zone.
caching nameserver makes a query for host1.foo.example.com to the nameserver for the **foo.example.com.** zone.
The **foo.example.com.** zone nameserver responds Authoritatively for the address **host1.foo.example.com**
The caching server stores all these queries and responses in its cache

### Query/Record Types

Different record types hold different info.
Specifying a quyer type ALL will fetch all record types.

#### A

Address Mapping Records
32bit IPv4 address

#### AAAA

IP version 6 Address Records
128Big ipv6 address

#### CNAME

Canonical Name Records
alias to another name

#### MX

Mail Exchanger Records
message transfer agents (mail servers) for a domain

#### NS

Nameserver Records
Delegate an Authoritative DNS zone nameserver

#### PTR

Reverse lookup pointer records
Pointer to a Canonical name (Ip address to name)

#### SOA

Start of authority Records
for a domain (domain and zone settings)

#### TXT

Text Records
Arbitrary human readable text, or machine readable data for specific purposes.

### Forward/Reverse DNS queries

#### Forward

use A or AAAA record types.
Most often used to turn a DNS name into an IP Address.
A FQDN is the full DNS Address  in the DNS DB, and the most significant part is the first (on the left)

#### Reverse

turn an ip address into a DNS name.
Uses PTR record type and arpa. domain in the DB.
in-addr.arps for IPv4
ip6.arpa for IPv6

In an IP address the most significant bit is on the right.
Ip address for PTR record in DATabase becomes:
191.168.13.32.in-addr.arps.

### DNS Server Daemons

Respons to info requests from clients.
* djbdns
    created by DJ Berstein
    very secure DNS alternative to BIND
* Dnsmasq
* Microsoft DNS Server
* BIND
    Berkeley Internet Domain Name
    de-factor standard for the internet

### DNS Client

The "Resolver"
DNS client / /etc/resolv.conf file nameserver entries were created at configuration time and never changed.
However, recently network configuration is becoming more dynamic and nameserver entries need to be adjusted.
Here are some Utilities to change the behaviour of /etc/resolv.conf:
* /etc/resolv.conf
    traditional static file used to configure the resolver.
* Network Interface Configuration Files
    * nameserver info may be added to i/f config files, overwriting or modifying /etc/resolv.conf when the i/f is started
* DHCP CLient
    * DHCP Server often provides nameserver info as part of info sent to DHCP client, replacing the existing nameserver records.
* resolvconf Service
    * most popular in Ubuntu
    * uses /etc/resolv.conf and a background service resolvconf.service to optimize the contents.
* dnsmasq
    * sets up in "mini" caching DNS server and may alter existing config to look at dnsmasq instead of /etc/resolv.conf
* systemd.resolved
    * DNS stub listener on IP address 127.0.0.53 on the loopback adapter and takes input from severl files. /etc/systemd/resolv.conf, /etc/systemd/network/\*.network and any info made available by dnsmasq

### BIND (named) Server

Widely-used, ISC standard DNS internet software. Available for most unix list systems.
BIND4 & BIND8 are obsolete, both have security problems.
BIND9 was a ground-up rewrite.
* Security issues fixed
* New features enabled:
  * can be run inside a chroot
  * DNSSEC (DNS Security Extensions)
  * allows for different views
  * IPv6 support
  * includes rndc (Remote Name Daemon Control) utility.

#### Configuration file

stored in /etc/named.conf or /etc/bind/named.conf

* listen-on
  * port and ip address to listen for connections
* listen-on-v6
  * IPv6 port and address to listen for connections
* allow-query                                                       
  * Controls hosts which can make queries to the server
* recursion
  * Controls the server acting as a recursive resolver
* forwarders
  * when acting as a recursive server, controls which nameservers we should query first.
* forward-first
  * Controls where the first recursive query happens

#### Zone Files

Authoritative zones are defined and must contain an SOA (Start of Authority) record and should contain an NS record.

SOA records are required for every zone.
Fields include:
* admin email
* primary name server
* serial number
* Timers: Refresh/Retry/Expire/Negative TTL

#### Split horizon

Also referred to as a DNS view.
can provide different DNS answers to requests depending on selection criteria.
uses multiple zone files with different data for the same zone.
often used to provide DNS services inside a corporation using private/non-routable addresses.

DNS view can assist solving challenges with multiple domains on a single DNS server:
* private network served by a DHCP service.
* ...

##### View configurations

Wrap the zone stanza with a view command.
Add to the view a name.
add to the view conditional-tests as required.
* match-clients
* match-connections
* recursive-only

### DNS Tools

#### dig

domain information groper.
Queries DNS for domain name or ip addres mapping.
output format resembles DNS records.
excellent for debugging DNS queries.

#### host

simple interface for DNS queries.
excellent for use in scripts.

#### nslookup

name server lookup.
saved from deprecation.
same as dig, but less verbose

#### nsupdate

name server update.
send updates to a name server.
requires authentication and permission.
uses dnssec.
