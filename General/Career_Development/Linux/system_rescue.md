# system rescue

inevitably systems will be trashed and require recovery.

## emergency boot media

useful when your system won't boot due to:

* missing, misconfigured or corrupted files / misconfigured service
* lost or scrambled root password

This is achieved through Live Media.
Which is a complete, bootable OS which runs in memory.
It can also be used to evaluate a new distro without installing it.

recovery/rescue mode is accessed from a menu option in the boot screen.
You're asked some question and then where to load the recuse media from.

many distros provide a boot.iso image for download. You can use dd to place it on a USB stick

### tools for creating emergency boot media

livecd-tools
livesub-creator

## emergency mode

booted into the most minimal environment possible

* root fs is read-only
* no init scripts run
* almost nothing setup

emergency mode is before single user mode. The usefulness of this is that if any init script is corrupted it won't be run

to enter emergency mode

* select an entry from the GRUB boot menu
* hit e, for edit
* add 'emergency' to the kernel parameter list
* the boot

## single user mode

if your system boots, but does not allow you to log in, try single user mode.

In single user mode:

* init is started
* services are started
* Network is not activated
* all possible filesystes are mounted
* root access is granted without password
* a system maintenance command line shell is launched

runlevel is 1

select an entry from the GRUB boot menu, hit e and add single to the kernel parameters