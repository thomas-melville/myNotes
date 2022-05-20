# RPM

Developed by RedHat
all files packaged into a single rpm file
rpm files also contain dependency information
    but not repository information

## +ives

determine what package a file is part of
determine what version is installed
install/uninstall packages
verify a package was installed correctly
distinguish docs from the rest of the package and decide whether to download them or not
use ftp / HTTP

## package file names

there is a standard, based on fields that represent specific info
binary:
<name>-<version>-<release>.<distro>.<architecture>.rpm
source
<name>-<version>-<release>.<distro>.src.rpm

## db directory

/var/lib/rpm is the default system directory
files in the form of Berkley DB hash files.
never manually edit the contents
--dbpath to change to another directory.
--rebuilddb, this is more for repair than periodic

## rpm programs

/usr/lib/rpm - scripts and programs

## installing a program

rpm performs a number of tasks when installing a program

* perform dependency checks
* perform conflict checks
* executes commands required before installation
* deals intelligently with congi files
* unpacks files from packages and installs them with correct attributes
* executes commands required after installation
* updates the system rpm db

## rpm command

-q              what version of package is installed
-qf             what package is the file in
-ql             what files are installed by this package
-qi             package information
-qip            package information from the file not the db
-qa             list all packages installed on the system
--requires      pre requisites for package
--whatprovides  what package provides a particular requisite version
-V              verify files from package are consistent with db

all the above flags can be mixed and matched

-i      install
-v      verbose (add it in multiple times to increase the logging level further)
-h      print hash marks while installing to show progress

-e      uninstall (erase) it fails if another package depends on this package
--test  test whether an uninstall will pass/fail

-U      upgrade or install a package / list of packages
--oldpackage    must add this to downgrade a package

-F      freshen packages (if a newer version exists, then upgrade. Otherwise do nothing)

## installing/upgrading the kernel

never upgrade the kernel with -U as it needs a reboot and you won't have the old version to roll back to
kernel versions can be installed side by side and you can decide which one to boot to.
GRUB file will be updated to include the new version

## rpm2cpio

copy files from the rpm to a cpio archive and extract them

## rpmbuild

create and manipulate source and binary packages

## rpm package structure

### spec file

Header section of key value pairs.
All identifiers are prepened with %
* %description - text description of the package
* %prep - shell script to unpack source code and apply patches
* %changelog
* ...
