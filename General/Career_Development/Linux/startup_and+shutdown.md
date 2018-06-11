# start up and shut down

## boot sequence

1. BIOS/UEFI locates and executes the boot program / loader
2. boot loader loads the kernel
3. kernel starts the init process, pid=1
4. init manages sys initialization, using systemd, Upstart or the older SysVinit startup scripts.

### 1. BIOS

Basic Input Output System
Contains all the code required to gain initial access to the hardware devices
Typically placed in the ROM chip.
When the power is applied to a system, it can only perform the operations the BIOS orders it to.
BIOS first runs POST, Power On Self Test.

* checks memory and hardware
* searches a specific location/device for a boot program

typically, the boot program is found in the devices MBR, Master Boot Record
Control of the system is then transferred to this boot program, usually GRUB

### 2. Boot loader

loads the kernel into memory and executes it
on many platforms the kernel first has to compress itself
it then performs the hardware checks
gains access to the important peripheral hardware
eventually runs the init process

### 3. init process

continues system start up
managing either systemd, Upstart, or running appropriate init scripts if SysVinit is used.

newer systems are moving to UEFI, a replacement for BIOS, which performs many of the same functions.

## Boot loaders

A number of different boot loaders in Linux:

* GRUB (GRand Unified Boot loader)
* LILO
* efilinux
* Das U-Boot

LILO, Linux Loader, is obsolete.

efilinux is an alternative boot loader specifically designed for the UEFI mechanism

Das U-Boot is the most popular loader for embedded linux systems

### GRUB

Virtually all, non-embedded, x86 modern Linux distro's use GRUB

GRUBs features:

* Boot multiple OSs
* alternative kernels can be chosen at boot time
* easily configure boot parameters
* both graphical and text based interface
* powerful CLI
* network based diskless booting

2 versions of grub 1 & 2. Philosophy is the same apart from some minor differences
My machine looks to be grub one, or is it UEFI????

At boot a basic config file is read from /boot/grub/grub.cfg or boot/grub2/grub.cfg
This file is auto generated based on the contents of /etc/grub.d directory and /etc/default/grub

after system boot and initial POST and BIOS stages the Boot loader menu is entered.
text / graphical view of possible images to boot

you can hit enter on any image to boot the image.

you can do more,

hit e to enter an interactive shell before booting the image
in this shell you can edit config files for the image
they changes aren't persisted, they are just for this boot of the image.

## shutdown command

bring the system down in a secure fashion
notify all users
stop it in a graceful and non-desctructive fashion

shutdown -h now