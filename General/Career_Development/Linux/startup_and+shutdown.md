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
multiple OSs and multiple kernel versions

you can do more,

hit e to enter an interactive shell before booting the image
in this shell you can edit config files for the image
they changes aren't persisted, they are just for this boot of the image.

## init

methods:

* System V Init
* upstart
* systemd!!

All major distros today use systemd

the init process is run by /sbin/init and is the first process on the system.
on my machine it is a link to /lib/systemd/systemd.
it runs until the system is shutdown.

* coordinates the later stages of the boot process
* configures all aspects of the environment
* starts processes needed for logging in

originally the init process was SysVInit.
very old and for different circumstances

* multi-user mainframe systems
* single processor system
* startup/shut down was not an important matter

To address the limitations of SysVInit and the change in the machine landscape new methods were developed.
The two main schemes are:

* Upstart
* * developed by Ubuntu
* * adopted in Fedora and RHEL 6
* systemd
* * more recent
* * Fedora was the first distro to offer it
* * Default on RHEL 7
* * this is the default and norm now

### systemd

features:

* compatible with SysVInit scripts, compatibility layer which allows systemd to consume sysvinit scripts
* boots faster than previous systems
* aggressive parallel capabilities
* Uses socket and d-bus activation for starting services
* replaces shell scripts with programs, .service files
* offers on demand starting of daemons
* keeps track of processes using cgroups
* supports creating snapshots and restoring of sys state
* maintains mount and automount points
* ...

all daemons are sorted into their own cgroups

configuration files:

* /etc/hostname         hostname
* /etc/vconsole.conf    default keyboard mapping and console font
* /etc/sysctl.d/*.conf  drop-in dir for kernel sysctl params
* /etc/os-release       distribution ID file

systemctl is the only utility you need for managing services

systemctl [options] command [name][.service]

most of the time you can omit .service

list, start, stop, status, enable, disable

enable/disable auto start the service on system start

### SysVinit

As a SysVinit system starts, it passes through a sequence of run levels which define different system states, 0 -> 6

2-5   define what services are running for a normal system, different distros define them differently

* 0     system halt state
* 1     single user mode
* 2     multiple user, no NFS, only test login
* 3     multiple user, with NFS and network, only text login
* 4     not used
* 5     multiple user, with NFS and n/w, graphical login with X
* 6     system reboot

you can see the current runlevel by executing **runlevel**

traditional method is to run the rc.sysinit script, in /etc and called rc.d.
there are a number of script rc[0-6].d which are run at each runlevel.
rc.local is for system specific apps.
all the actual scripts are in /etc/init.d.
everything in rc is ac symbolic link.
Start scripts start with S
Kill scripts start with K
The number after K/S determines the order the scripts are executed in
**chkconfig** should be used to manage the links which will control what gets executed at each runlevel

### chkconfig

query and configure what runlevels the various system services are run in.

You can add your own services and write your own start up scripts.
They can be added by chkconfig then

### services

every OS has services which are started on sys init and rename until shutdown.
you can start, stop, status them anytime.
you can use the **service** command

sudo service <service> <command>

commands differ depending on the service specified

all service really does is cd to /etc/init.d and run the appropriate script with the supplied options

### Upstart

is event driven
notifications are sent to the init process to tell it to execute certain commands
superseded by systemd

rcS.conf script is executed after kernel starts init process
this causes rc-sysinit.conf to be run
rc-sysinit.conf does a number of things

* starts LVM
* mounts FSs
* executes the run level scripts that are specified in /etc/inittab

the command for upstart is **initctl**, you can view, start, stop jobs in much the same way as service

## shutdown command

bring the system down in a secure fashion
notify all users
stop it in a graceful and non-desctructive fashion

shutdown -h now