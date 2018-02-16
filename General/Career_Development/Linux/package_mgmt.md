# Package Management

installing and removing software
guarding against corruption
verify authenticity
upgrading and updating

a package may contain exe files, data files, documentation, install scripts & config files, and metadata

All info is stored locally in an internal database

## Features

Automation: no need for manual installs and upgrades
Scalability: install package on 1 / 1,000 systems
Repeatability:
Security

## families

2 level of utilities:

* low level (installing, verifying, removing individual packages or groups of packages)
* higher level (maintain dependencies between packages, all required software is there when a package is installed, and removed when no longer needed)

### RPM (Redhat)

RPM (Redhat Package Manager)
distributions: Red hat, centos, fedora, suse

low level:  rpm
high level: yum, dnf or zypper

### APT (Debian)

APT (Advanced packing Tool)
distributions: Debian, Ubuntu, Mint

low level:  dpkg
high level: apt-*

### other package management systems

portage/emerge used by Gentoo
pacman used by Arch
specialized ones used by embedded linux systems and Android.
tarballs in Slackware

## Package Types

Binary

* contains files ready for deployment, exe & lib files.
* architecture dependent
* compiled for each machine type

Source

* used to generate binary packages
* one source package can be used for multiple architectures

Architecture-independent

* run under script interpreters
* docs and config files

Meta-packages

* groups of associated packages that collect everything needed to install a relatively large sub-system.

## Package Sources

every distribution has one or more package repositories
system utils go here to obtain/update software.
you can add more, but than can lead to dependency hell

## Building your own software package

this allows you to control what goes in and how its used.
you can create the package so that installing it runs scripts that perform tasks:

* create symbolic links
* create directories
* set permissions
* anything that can be scripted