# network configuration

## Layer 2

Common options to change are MTU (Max Transmission Unit), link speed, link status

some tools used:

ip
  change mtu

ethtool
  change link speed

mii-tool
  check link status

for changing additional parameters there are two options:

kernel modules modinfo and modprobe using /etc/modprobe.d
udev using the drop-in directory /etc/udev/rules.d

udev = user device facility
  used to manage network hardware interfaces for the linux kernel
  rename interfaces or change configs to match MAC addresses

udevadm command is used for control, querying and debugging udev config files.

Additional layer 2 addresses maybe added to an adapter through the use of MACVlan (uses network namespaces)

MACVlan modes
* Private
    instances can't talk to each other.
    only external communication
* VEPA
    Virtual Ethernet Port Aggregator
    instances can talk to each other through external hardware
* Bridge
    instances are allowed internal and external communication
* Passthru
    instances are directly connected to the interface
* Source
    employs MAC filtering

## Layer 3

Common things to change are network address, default route, add DNS server, manually request DHCP config

some tools used:

ip

route

echo >> /etc/resolv.conf

dhclient


network settings are stored in configuration files so they are persisted across reboots.
Network Manager is an example of a config tool.
systemd-network relies on these config files.
netplan is another option to create nw config files.
creates confif files at runtime from pre-defined yaml files.
  files located in /etc/netplan folder
  two important commands try and apply
  try calls apply, but will rollback if apply fails
  on startup generate command is called,
    builds network config files from the plans and stores them where Network Manager can find them.
Find out which tool is configuragin your network:
Network Manager: nmcli
systemd-network: networkcli

## Network Manager

Part of most popular Linux distros. each distro can write their own plugins/interface.

It provides:
* GUI tool
    nm-connection-editor
* applet
    nm-applet
* text interface
    nmtui
* cli interface
    nmcli

If a network interface has no network config file then Network Manager will configure it using DHCP

Adapters:
* Bluetooth
* DSL/PPoE
* Ethernet
* InfiniBand
* Mobile Broadband
* Wi-Fi

Virtual:
* Bond
* Bridge
* IP Tunnel
* MACSec
* Team
* Vlan
* Wireguard

## Distro specific configurations

### OpenSUSE

located in /etc/sysconfig/network and match ifcfg-<interface> pattern.
hostname is configured in /etc/HOSTNAME.
DNS client settings are managed by either:
* /etc/sysconfig/network/config file
* manually editing /etc/resolve.conf
Managed by a tool called netconfig

### Ubuntu

network config files located in /etc/network directory
interface config files located in /etc/network/interfaces directory.
hostname is configured in /etc/hostname.
DNS configuration is managed by the resolvconf tool and nameserver configuration is managed in /etc/network/interfaces.

## VPN

Virtual Private Network. Provide a secure connection through an unsecure network. Data is encrypted to avoid unwanted exposure.
OpenVPN and WireGuard are two VPN Software providers.
Orgs take advantage of VPN to make it easy for remote workers to connect to:
* Printers
* File shares
* databases
* security
VPN types are provided by different protocols
* Secure Transport Layer (SSL/TLS)
    OpenVPN
* Internet Protocol Security (IPSec)
* Datagram Transport Layer (DTLS)
* Secure Shell (SSH) VPN

OpenVPN
* Can traverse NAT and Firewalls
* Can run over TCP/UDP
* cross-platform
