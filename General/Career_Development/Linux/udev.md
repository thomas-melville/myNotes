# udev

intelligent apparatus for discovering hardware and peripheral devices both during system boot up and later when they are connected to the system.

udev is a subsystem which handles new devices when they are plugged in

all devices are placed in the /dev folder.

major and minor numbers identify the driver associated with a device. Each driver uniquely reserves a group of numbers

in ls of dev the major and minor numbers are where the file size is normally displayed.

* the major number is the driver number
* the minor number is the device identifier.

udev creates device nodes on the fly as they are needed.
the u stands for user

* runs as a daemon, systemd-udevd on my machine
* monitors the netlink socket
* uevent kernel facility sends a message through the socket
* udev receives this
* takes appropriate action to create/remove device nodes according to the rules.

3 components:

* libudev       allows access to information about the devices
* systemd-udevd daemon that manages the /dev directory
* udevadm       utility for control and diagnostic

udev works with the hotplug subsystem.

information is gathered from /sys after the add/remove event is triggered.
and a set of config files.

there are also rules which are parsed and can be found in /etc/udev/rules/*.rules