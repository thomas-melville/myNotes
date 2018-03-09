# Linux Directory Structure

## FHS

FileSystem Hierarchy Standard
https://refspecs.linuxfoundation.org/FHS_3.0/fhs-3.0.pdf


Dir			Description
/					The directory called "root." It is the starting point for the file system hierarchy. Note that this is not related to the root, or superuser, account.
/bin			Binaries and other executable programs. No subdirectories, just all executable files.
                All binaries required by root to boot must be here.

/boot			Files needed to boot the operating system.
                essential files are: vmlinux-<version> (compressed kernel) and initrd (initial RAM file system, which is used before the real file system is available)
                other important files are config (config file for when compiling the kernel) and System.map (hex addresses of all kernel symbols)

/cgroup			Control Groups hierarchy.

/dev			Device files, typically controlled by the operating system and the system administrators.
                represent devices built into or connected to the system and are essential for the system to operate

/etc			System (machine local) configuration files and some startup scripts.
                skel - skeleton for creating new users home directories
                init.d - start up / shutdown scripts when using System V initilization

/home			Home directories of individual users and groups of users

/lib			System Libraries needed to execute the binaries in bin and sbin. 
                Important for booting the system and executing commands on the root FS
                Also contains kernel modules under /lib/modules/<kernel-version>
                PAM (Pluggable Authentication Modules) are located under lib/security.
                for 32 bit

/lib64			System Libraries, 64 bit.
                Not everyone follows this though, you will find 64 bit stuff in lib

/lost+found		Used by the file system to store recovered files after a file system check has been performed.

/media			Used to mount removable media like CD-ROMs.
                modern linux systems mount them automtically upon insertion.
                udev creates a directory under media

/mnt			Used to mount external file systems.
                NFS, Samba, CIFS, AFS

/opt			Optional or third party software.
                Keeps all their files in one isolated place instead of spread through out the FS with the other software

/proc			Provides information about running processes.
                mount point for a pseudo file system, where all information only resides in memory, not on disk.
                kernel exposes important data structures through proc entries.
                each active process has it's own subdirectory with the PID as the directory name
                entries in proc are often called virtual files, listed as having 0 bytes but contain information.
                interrupts, meminfo, mounts, partitions give an up to date glimpse of the hardware
                information is gathered when the files are inspected, 

/root			The home directory for the root account.

/run            store transient files that contains runtime information
                pseudo FS

/sbin			System administration binaries
                booting, restoring, recovering, repairing

/srv			Contains site-specific data which is served by the system. such folders underneath can be www, ftp. Rarely used
                makes it easy for admins to find files for a particular service
                similar purpose to var

/sys			Used to display and sometimes configure the devices and buses known to the Linux kernel.
                another pseudo file system
                gather system info and modify its behaviour while running.

/tmp			Temporary space, typically cleared on reboot, this is the Ubuntu policy.
                Other distributions clear files older than x days
                avoid putting large files in tmp, can slow down / crash the system

/usr			User related programs.
                secondary hierarchy, files not required by booting
                no sub directories by packages directly under usr, symbolic links if required.,
                read only files, binaries, ...
                structure under user is similar to structre under /
                    bin, etc, lib, lib64, sbin, tmp, ...
                man pages are stored here

/usr/bin		Binaries and other executable programs.
                which are not required by root to boot

/usr/lib		Libraries.

/usr/local		Locally installed software that is not part of the base operating system.

/usr/sbin		System administration binaries.

/var			Variable/volatile data, most notably log files.


applications can follow a directory structure convention:
bin
etc
lib
logs

or they can be split across the corresponding folders in the / directory

## Partitioning

root directory is some times in it's own partition with user specific directories in other partitions

## taxonomy of information

### Shareable V Non-Shareable

specific to a host or can be used by many hosts

### Variable V Static

anything that requires system admin assistance to change is static