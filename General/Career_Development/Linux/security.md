# security

number of different ideas on how to incorporate a computer security module into linux.
The idea is to implement mandatory access control over the variety of requests made to the kernel,
but do so in a way that:

1. minimizes changes to the kernel
2. minimizes overhead on the kernel
3. Permits flexibility and choice between diff impls, each of which is presented as a self-contained Linux Security Module (LSM)

hook system calls, which is to insert code whenever an app requests a transition to kernel mode
this code makes sure:

* permissions are valid
* malicious intent is protected against

Current LSM implementations are: (in order of usage volume)

* SELinux
* AppArmor
* Smack
* TOMOYO

only one can be used at a time

## SELinux

originally developed by the NSA
integral to RHEL for a very long time
Operationally a set of security rules that are used to determine which processes can access which files, directories, ports and other items

3 conceptual quantities:

* Contexts
  * labels to files, processes and ports.
  * ex: user, role, type
* Rules
  * describe access control in terms of contexts, processes, files, ...
* Policies
  * a set of rules that describe what system wide access control decisions are made by SELinux

Default policy is to deny all access
rules are used to describe allowed actions on the system

Many standard command lines tools, ls/ps/cp/mv/mkdir were extended to support SELinux
often the Z parameter is passed

### Modes

* Enforcing
  * All SeLinux code is operative and access is denied according to policy
  * all violations are audited and logged
* Permissive
  * Enables SELinux code but only warns when access would be denied when enforced.
  * everything audited and logged
* Disabled
  * completely disable all the SELinux code

configuration is stored in /etc/selinux/config

There are two ways to enable SELinux:

* configuration file
  * edit /etc/selinux/config
* kernel parameter
  * add selinux=<x> to en/disable it

### commands

**sestatus** see the current status of the SELinux setup.

* currently disabled on my machine

**getenforce** examine current mode

**setenforce** set current mode between Enforcing and Permissive on the fly

* can't change from/to disabled with setenforce

**restorecon** resets file contexts, based on parent directory setting

**semanage fcontext** change and display the default context of files and directories

### Policies

stored in same config file.
multiple policies allowed in the file, but only one active.
changing the policy may require a reboot and a time consuming re-labelling of FS contents.
Each policy has files that must be installed under /etc/selinx/[type]
Currently there is only default on my machine

Most common are:

* targeted
  * default policy
  * in which SELinux is more restricted to targeted processes.
  * User processes and init processes are nor targeted.
  * memory restrictions for all processes, which reduces the vulnerability to buffer overflow attacks.
* minimum
  * a modification of the targeted policy where only selected processes are protected.
* MLS
  * Multi-Level Security policy is much more restrictive
  * all processes are placed in fine-grained security domains with particular policies

#### SELinux Booleans

allows policy behaviour to be changed at runtime without rewriting policy
configure these parameters to be enabled / disabled.

**getseboole** & **setseBool** to display and change the value

### Contexts

labels applied to files, directories, ports and processes.
describe access rules.

Four Contexts:

* User
* Role
* Type
  * most common
* Level

naming context is that it ends in a letter which describes its type.

newly created files inherit the context from the parent directory, but if a file is moved the context stays the same.
The classic issue is moving files to the DocumentRoot of a httpd server.
On SELinux-enabled system , the web server can only access files with the correct context labels.
Creating a file anywhere else and moving it into the web folder will be result in it not being accessible to the web server.

### Trouble shooting

the setroubleshoot-server package has commands which help with troubleshooting

**sealert** is one of them

**audit2allow** generates policy rules from denied operations

**audit2why** translates audit messages into descriptions of why the access was denied

## AppArmor

**AppArmor is running ony my laptop.**

LSM alternative to SELinux
support incorporated into the kernel since 2006
used by SUSE, Ubuntu & others

* Provides (Mandatory Access Control) MAC
* associate a security profile to a program which restricts its capabilities
* some consider it easier than SELinux
* considered file system neutral

Supplements UNIX Discretionary Access Control (DAC) model by providing MAC

Includes a learning mode, where violations are logged but not prevented.
This log can be turned into a profile.

apparmor is a service so you can use **systemctl** to check it's status

### commands

**apparmor_status** to see the current status

profiles and processes are either in enforce or complain mode.
directly analogous to enforcing and permissive in SELinux

**aa-enforce** to change the mode to enforce for a profile

**aa-complain** to change the mode to complain for a profile

... many more

### Profiles

restrict how executable programs can be used

Two modes:

* Enforce
  * apps are prevented in acting ways which are restricted.
  * violations are logged
  * this is the default mode
* Complain
  * policies not enforced
  * violations are logged
  * this is also called learning mode

Distros come with prepackaged profiles, which are installed by packages and others by arrarmor itself using **apparmor-profiles**

## Local System Security

no computer can every be absolutely secure.
all we can do is slow down/discourage attackers.
Security can be defined in terms of the systems ability to regularly do what it is supposed to do, integrity and correctness of the system, and ensuring that the system is only available to those authorized to use it.
The biggest problem with security is finding the appropriate mix of security and productivity.
The human factor is the weakest link in the security chain.

4 areas:

* physical
* local
* remote
* personnel

### security philosophies

* anything not expressly permitted is denied
* anything not expressly forbidden is permitted

the first one is the most commonly used one
the second assumes a lot of trust

### advice

* never put the current directory in your $PATH variable
  * this is because a user could create a script with the same name as a common linux command with malicious intent and someone could unwittingly execute it
* protect your BIOS with a password
* secure the boot process with a password too
* use the security features when mounting