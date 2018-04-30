# virtualization

virtualization can be done for more than just machines:

* network
* storage
* application

implemented originally decades ago on mainframes:

* enables better hardware utilization
* OS's often progress faster than hardware
* it is microcode driven
* not particularly user friendly.

later, it migrated to PCs/desktops

* initially done through emulation
* CPUs enhanced to support virtualization led to a boost in performance, easier config, more flexibility in VM installation and migration.

## terms

* host      underlying physical OS managing one or more VMS
* guest     the vm, which is an instance of a complete OS

## emulators

the first impl of virtualization on the PC arch was through the use of emulators
An application running on the current OS would appear to another OS as a specific hardware env
don't require special hw to operate
Qemu is one such emulator
runs completely in software
hw constructs are replaced by software
useful for running vms on different architectures
performance is relatively slow

## hypervisors

also known as Virtual Machine Monitors (VMM)
a hypervisor initiates, terminates and manages guests.
essentially it's the host system

2 types:

* hardware virtualization (full virtualization)
* * guest system runs without being aware that it is virtualized
* * doesn't require modifications to run in this fashion
* Para-virtualization
* * guest system is aware it is virtualized

CPUs now incorporate virtualization extensions to the architecture that allow hypervisor to run fully virtualized guest systems with only minor performance hits
look for the vmx flag in /proc/cpuinfo to see if your pc supports hardware virtualization

The hypervisor can be:

* External to the host OS kernel: VMWare
* Internal to the host OS kernel: KVM

Going past emulators, the merging of the hypervisor into a specially designed lightweight kernel was the next step
This is VMWare

The KVM project added hypervisor capabilities to the linux kernel.
enabled the kernel itself to become a hypervisor

### list of hypervisors

QEMU/KVM
Xen
VirtualBox
VMWare
Microsoft Hyper-V
LXC

## libvirt

toolkit to interact with virtualization technologies.
provides management for virtual machines, networks and storage
many application programs interface with libvirt.
examples: virt-manager, virt-viewer, virt-install, **virsh**

## QEMU

Quick EMUlator
performs hw emulation, or virtualization
emulates CPUs by dynamically translating binary instructions between the host arch and the emulated one
using it KVM you can reach speeds close to those of the native host.
supports many architectures
you can use it with other hypervisors too
VirtualBox, use qcow2 formatted images

### disk image file formats

supports many but only 2 used primarily

* raw       default, simplest and easiest to format
* qcow2COW  Copy On Write

others

* vdi       Virtual box
* vmdk      VMWare

## kvm

uses linux kernel for computing resources
engages in a co-processing relationship with the linux kernel.

In this format, kvm runs the virtual machine monitor (the hypervisor) within one or more of the CPUs, using VMX or SVM instructions
At the same time the kernel is executing on the other CPUs.
VMM runs the guest, which is running at full hardware speed, until it executes an instruction that causes the VMM to take over
VMM can use any Linux resource to emulate a guest instruction

### commands

virt-*& qemu-* are the command line utilities for managing virtual machines images
virt-manager, kimchi, OpenStack are graphical interfaces for the same.