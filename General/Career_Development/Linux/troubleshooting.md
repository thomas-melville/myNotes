# troubleshooting

three levels of troubleshooting

* beginner
  * can be taught very quickly
* experienced
  * after years of a experience
* wizard
  * the one

## basic techniques

take a number of steps iteratively until a solution is found.

1. Characterize the problem
2. Reproduce the problem
3. Always try the easiest things first
4. Eliminate possible causes one at a time
5. Change only one thing at a time. If it doesn't fix the problem change it back
6. Check the system logs

## networking

things to check:

* IP Configuration
  * ip
  * ifconfig (only older systems)
* Network Driver
  * lsmod, to check if the network driver is loaded
  * relevant pseudo-files
* Connectivity
  * ping
  * traceroute, follow packets in a single instance
  * mtr, follow packets in a continuous fashion
* Default gateway and routing configuration
  * route -n
* hostname resolution
  * dig
  * host

## file integrity

packaging systems have methods for verifying file integrity and checking for changes

* rpm -V <package> (RPM)
* debsums options <package> (Debian)
* dpkg -V <package>

## boot process failures

Assuming you get to the BIOS stage, you may end up in one of the following states

* No boot loader screen
  * check GRUB for misconfiguration
  * check for a corrupt boot sector.
  * you may have to re-install the boot loader
* Kernel fails to load
  * most likely misconfigured, corrupt kernel, or incorrect parameters
  * check what changed????
  * reinstall kernel
  * or go into interactive GRUB and start with minimal parameters
  * or boot into a rescue image
* Kernel loads, but fails to mount root filesystem
  * misconfigured GRUB config file
  * misconfigured /etc/fstab
  * no support for the boot filesystem type
  * fsck can be used to repair if fstab isn't misconfigured
* failure during the init process
  * examine the logs closely, as there are many things that could go wrong
  * try booting into a lower run level

## rescue installed system

you should be able to boot from this.
contains a limited but powerful set of utilities

* Disk Maintenance utilities
  * fdisk
  * mdadm
  * pvcreate
  * vgcreate
* Networking utilities
  * ifconfig
  * route
  * traceroute
* Miscellaneous utilities
  * bash
  * chroot
  * ...
* Logging utilities