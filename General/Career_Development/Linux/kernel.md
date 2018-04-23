# kernel

Narrowly defined, Linux is only the kernel of the operating system.
provides the interaction between the hardware and the software.
handles all connected devices using device drivers.

## main responsibilities

* system initialization and boot up
* process scheduling
* memory management
* controlling access to hardware
* i/o between applications and storage devices
* impl of local and network file systems
* security control, both locally and over the network
* networking control

system is booted on kernel command line
config stored in /boot/grub/grub.cfg
to see what command the kernel was started with: cat /proc/cmdline
long list of parameters available. (man bootparam)

some examples

* root              root fs
* ro                mounts root device read-only on boot
* vconsole.keymap   which keyboard to use on the console
* crashkernel       how much memory to set aside for kernel crashdumps
* rhgb              for graphical boot
* quiet             disables most logging messages

## modules

Many facilities in the kernel are designed to be built-in when the kernel is initially loaded, or added/removed later as modules.
All but the most central kernel components are modules.

* device drivers
* networking
* filesystems (mkfs loads the required fs module)

also, facilitates development as a kernel reboot is not needed to load in a module.

even with the widespread usage of modules, the kernel is still a monolithic kernel architecture. Once a module is loaded it becomes a fully fledged part of the kernel.

normal fs location for modules is /lib/modules/<kernel.version>/.../...
a kernel module always has the extension ko
modules are kernel version specific

in most cases modprobe is used instead of insmod and rmmod

* add -r for removing the specified module

if a module is being used by another module it is impossible to unload it.
if a module is being used by a process it is impossible to unload it.
when a module is loaded by modprobe any dependencies it has will be loaded first.
when a module is unloaded by modprobe it will remove all dependencies which aren't being used by another module/process.

### commands

* lsmod     list loaded modules
            format: name    size    processes using module  other modules using module
* insmod    load a module (must be done as root) and file extension is required.
            can take parameters which are passed into module
* rmmod     remove a module, file extension is not required
* modprobe  load/unload modules, using a pre built module db with dep info
            can take parameters which are passed into module
* depmod    rebuild the module dependency db
* modinfo   display info about a module

the same information as modinfo can be seen in the sys pseudo file system.

/sys/module/<module_name>

Files in the /etc/modprobe.d/ directory control some parameters that come into play when loading modules. module name aliases, module blacklist

## sysctl

used to read and tune kernel parameters at runtime
each value corresponds to a particular pseudo file residing under /proc/sys
    directory slashes being replaced by dots.
the sysctl function can be called from programs

systemd is a successor to sysctl

### arguments

-a      displays all current values
-p      read in and apply all parameters set in /etc/sysctl.conf
